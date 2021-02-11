/*
Copyright © 2021 Abdul Fattah Ikhsan <ikhsannetwork@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ikhsanalatsary/MeowTube/instances"
	"github.com/ikhsanalatsary/MeowTube/interfaces"
	"github.com/ikhsanalatsary/MeowTube/logger"
	"github.com/ikhsanalatsary/MeowTube/vlc"
	"github.com/spf13/cobra"
)

var audioOnly bool
var fullscreen bool
var resolution string
var resolutions = map[string]string{
	"144p":    "144p",
	"240p":    "240p",
	"360p":    "360p",
	"480p":    "480p",
	"720p":    "720p",
	"720p60":  "720p60",
	"1080p":   "1080p",
	"1080p60": "1080p60",
}

var audioFormat string
var videoFormat string

// var encodingFormat string

// Audio format
// - m4a
// - webm

// Video format
// - mp4
// - webm
// -

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play YoutubeURL",
	Short: "To play YouTube video",
	Long: `This command requires videoID or youtube url without options.
The difference with video commnad is, this command not supported any options`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		videoID := args[0]
		if interfaces.IsValidYoutubeURL(args[0]) {
			videoID = interfaces.GetVideoIdFrom(args[0])
		}
		videoCmd.Run(cmd, []string{videoID})
	},
}

var videoCmd = &cobra.Command{
	Use:   "video videoId",
	Short: "To play YouTube video",
	Long:  `This command requires videoID or youtube url with optional options`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Request video...")
		videoID := args[0]
		if interfaces.IsValidYoutubeURL(args[0]) {
			videoID = interfaces.GetVideoIdFrom(args[0])
		}
		detailURL := "/api/v1/videos/" + videoID + "?fields=formatStreams,title,author,genre,adaptiveFormats,lengthSeconds"
		source := instances.FindFastest(detailURL)
		if source.Error != nil {
			logger.ThrowError(source.Error)
		}
		// fmt.Println("Requesting " + source.FastestURL)
		defer source.Resp.Body.Close()
		data, err := ioutil.ReadAll(source.Resp.Body)
		res, err := interfaces.UnmarshalFormatStream(data)
		if err != nil {
			logger.ThrowError(err)
		}
		flags := []string{
			"--network-caching=1000",
			"--video-title=" + res.Title,
			"--meta-title=" + res.Title,
			"--meta-artist=" + res.Author,
			"--meta-author=" + res.Author,
			"--meta-genre=" + res.Genre,
			"--input-title-format=" + res.Title,
			"--duration=" + fmt.Sprint(res.LengthSeconds),
			res.FormatStreams[0].URL,
		}
		if resolution != "" {
			var message string
			if len(res.AdaptiveFormats) > 1 {
				if _, ok := resolutions[string(resolution)]; !ok {
					logger.ThrowError("Invalid resolution")
				}
				for _, v := range res.AdaptiveFormats {
					if v.Container != nil && string(*v.Container) == videoFormat && v.QualityLabel != nil && string(*v.QualityLabel) == resolution {
						fmt.Println("resolution = ", *v.QualityLabel)
						flags[len(flags)-1] = v.URL
						message = ""
						break
					} else {
						message = "Unable to find resolution. Default resolution used as a fallback!"
					}
				}
				if message == "" {
					flags = append(flags, ":input-slave="+res.AdaptiveFormats[0].URL, ":network-caching=1000")
				} else {
					fmt.Println(message)
				}
			}
		}
		if fullscreen {
			// println("fullscreen")
			flags = append(flags, "--fullscreen")
		}
		// fmt.Println(strings.Join(flags, " "))
		VLC.Execute(flags...)
	},
}

var audioCmd = &cobra.Command{
	Use:   "audio :videoId",
	Short: "To play audio only",
	Long:  `This command requires videoID or youtube url without options`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Request audio...")
		videoID := args[0]
		if interfaces.IsValidYoutubeURL(args[0]) {
			videoID = interfaces.GetVideoIdFrom(args[0])
		}
		detailURL := "/api/v1/videos/" + videoID + "?fields=formatStreams,title,author,genre,adaptiveFormats,lengthSeconds"
		source := instances.FindFastest(detailURL)
		if source.Error != nil {
			logger.ThrowError(source.Error)
		}
		// fmt.Println("Requesting " + source.FastestURL)
		defer source.Resp.Body.Close()
		data, err := ioutil.ReadAll(source.Resp.Body)
		res, err := interfaces.UnmarshalFormatStream(data)
		if err != nil {
			logger.ThrowError(err)
		}
		flags := []string{
			"--network-caching=1000",
			"--video-title=" + res.Title,
			"--meta-title=" + res.Title,
			"--meta-artist=" + res.Author,
			"--meta-author=" + res.Author,
			"--meta-genre=" + res.Genre,
			"--input-title-format=" + res.Title,
			"--duration=" + fmt.Sprint(res.LengthSeconds),
		}
		if len(res.AdaptiveFormats) > 1 {
			for _, v := range res.AdaptiveFormats {
				if strings.Contains(v.Type, "audio") && string(*v.Container) == audioFormat {
					flags = append(flags, v.URL)
					break
				}
			}
		} else {
			logger.ThrowError("Cannot play stream")
		}
		VLC.Execute(flags...)
	},
}

func runPlaylist(cmd *cobra.Command, args []string) {
	fmt.Println("Request playlists...")
	playlistURL := "/api/v1/playlists/" + args[0]
	source := instances.FindFastest(playlistURL)
	if source.Error != nil {
		logger.ThrowError(source.Error)
	}
	defer source.Resp.Body.Close()
	data, err := ioutil.ReadAll(source.Resp.Body)
	res, err := interfaces.UnmarshalPlaylist(data)
	if err != nil {
		logger.ThrowError(err)
	}
	pl := &vlc.VLCPlaylist{}
	pl.Xmlns = "http://xspf.org/ns/0/"
	pl.Text = "xmlns"
	pl.Vlc = "http://www.videolan.org/vlc/playlist/ns/0/"
	pl.Version = "1"
	pl.Title = "Playlist"
	Tracks := []vlc.Track{}
	Items := []vlc.ExtensionItem{}
	if len(res.Videos) > 0 {
		playlists := instances.RequestAllPlaylist(source.FastestURL, res.Videos)
		if len(playlists) == 0 {
			logger.ThrowError("Requested videos not available!")
		}
		fmt.Println("Total videos: ", len(playlists))
		flags := []string{
			"--network-caching=1000",
		}
		for i, v := range playlists {
			id := fmt.Sprint(i)
			if v != nil {
				localOption := []string{
					"video-title=" + v.Title,
					"input-title-format=" + v.Title,
					"meta-title=" + v.Title,
					"meta-artist=" + v.Author,
					"meta-author=" + v.Author,
				}
				if audioOnly {
					if len(v.AdaptiveFormats) > 1 {
						for _, a := range v.AdaptiveFormats {
							if strings.Contains(a.Type, "audio") && string(*a.Container) == audioFormat {
								trEx := vlc.TrackExtension{
									Application: "http://www.videolan.org/vlc/playlist/0",
									ID:          id,
									Option:      localOption,
								}
								tr := vlc.Track{
									Location:  a.URL,
									Extension: trEx,
									Creator:   v.Author,
									Title:     v.Title,
									Duration:  fmt.Sprint(v.LengthSeconds),
								}
								Tracks = append(Tracks, tr)
								exItem := vlc.ExtensionItem{
									Tid: id,
								}
								Items = append(Items, exItem)
								break
								// flags = append(flags, a.URL, ":video-title="+v.Title, ":meta-title="+v.Title, ":meta-artist="+v.Author, ":meta-author="+v.Title)
							}
						}
					}
				} else {
					if len(v.FormatStreams) > 0 {
						trEx := vlc.TrackExtension{
							Application: "http://www.videolan.org/vlc/playlist/0",
							ID:          id,
							Option:      localOption,
						}
						tr := vlc.Track{
							Location:  v.FormatStreams[0].URL,
							Extension: trEx,
							Creator:   v.Author,
							Title:     v.Title,
							Duration:  fmt.Sprint(v.LengthSeconds),
						}
						Tracks = append(Tracks, tr)
						exItem := vlc.ExtensionItem{
							Tid: id,
						}
						Items = append(Items, exItem)
						// flags = append(flags, v.FormatStreams[0].URL, ":video-title="+v.Title, ":meta-title="+v.Title, ":meta-artist="+v.Author, ":meta-author="+v.Title)
					}
				}
			} else {
				continue
			}
		}
		pl.TrackList = vlc.TrackList{
			Track: Tracks,
		}
		pl.Extension = vlc.Extension{
			Application: "http://www.videolan.org/vlc/playlist/0",
			Item:        Items,
		}
		tmpFile, err := ioutil.TempFile(os.TempDir(), "playlist-"+"*.xspf")
		if err != nil {
			logger.ThrowError("Cannot create temporary file", err)
		}

		fmt.Println("Created Temporary Playlist File: " + tmpFile.Name())

		// Example writing to the file
		text, err := vlc.MarshalFrom(pl)
		if err != nil {
			logger.ThrowError("Failed to marshal from data", err)
		}
		if _, err = tmpFile.Write(text); err != nil {
			logger.ThrowError("Failed to write to temporary file", err)
		}
		flags = append(flags, tmpFile.Name())
		// Remember to clean up the file afterwards
		defer os.Remove(tmpFile.Name())
		VLC.Execute(flags...)
		fmt.Println("Deleting Temporary Playlist")

		// Close the file
		if err := tmpFile.Close(); err != nil {
			logger.ThrowError(err)
		}
	} else {
		logger.ThrowError("No videos found!")
	}
}

var playlistCmd = &cobra.Command{
	Use:   "playlist :playlistId",
	Short: "To play all videos from YouTube playlist",
	Args:  cobra.MaximumNArgs(1),
	Run:   runPlaylist,
}

var listCmd = &cobra.Command{
	Use:   "list :playlistId",
	Short: "Shorthand for playlist. To play all videos from YouTube playlist",
	Args:  cobra.MaximumNArgs(1),
	Run:   runPlaylist,
}

func init() {
	playCmd.AddCommand(videoCmd, audioCmd, playlistCmd, listCmd)
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	videoCmd.Flags().BoolVarP(&fullscreen, "fullscreen", "f", false, "Fullscreen video output (default \"disabled\")")
	videoCmd.Flags().StringVarP(&resolution, "resolution", "r", "", "Select high resolution streaming 144p, 240p, 360p, 480p, 720p, 720p60, 1080p, 1080p60  (default \"360p\")")
	videoCmd.Flags().StringVarP(&videoFormat, "format", "F", string(interfaces.Mp4), "Select video format streaming mp4, webm (default mp4)")
	playlistCmd.Flags().BoolVarP(&audioOnly, "audio-only", "a", false, "Play the playlist in audio format only")
	playlistCmd.Flags().StringVarP(&audioFormat, "format", "F", string(interfaces.M4A), "Select audio format streaming m4a, webm")
	listCmd.Flags().BoolVarP(&audioOnly, "audio-only", "a", false, "Play the playlist in audio format only")
	listCmd.Flags().StringVarP(&audioFormat, "format", "F", string(interfaces.M4A), "Select audio format streaming m4a, webm")
	audioCmd.Flags().StringVarP(&audioFormat, "format", "F", string(interfaces.M4A), "Select audio format streaming m4a, webm")
}
