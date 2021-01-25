/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"os"
	"strings"

	"github.com/ikhsanalatsary/MeowTube/instances"
	"github.com/ikhsanalatsary/MeowTube/interfaces"
	"github.com/ikhsanalatsary/MeowTube/vlc"
	"github.com/spf13/cobra"
)

var audioOnly bool
var fullscreen bool
var resolution string
var resolutions = map[string]string{
	"144p":  "144p",
	"240p":  "240p",
	"360p":  "360p",
	"480p":  "480p",
	"720p":  "720p",
	"1080p": "1080p",
}

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("VLC => ", VLC.GetVlc())
		if len(args) > 0 {
			if len(args) == 1 {
				fmt.Println("No Command")
			} else if len(args) == 2 {
				// if args[0] == "video" {
				// 	// detailURL := source.FastestURL + "/api/v1/videos/" + args[1] + "?fields=formatStreams,title,author,genre"
				// 	detailURL := "/api/v1/videos/" + args[1] + "?fields=formatStreams,title,author,genre,adaptiveFormats"
				// 	source, err := instances.FindFastest(&instances.InstanceList, detailURL)
				// 	if err != nil {
				// 		log.Fatal(err)
				// 	}
				// 	// fmt.Println("Requesting " + source.FastestURL)
				// 	defer source.Resp.Body.Close()
				// 	data, err := ioutil.ReadAll(source.Resp.Body)
				// 	res, err := interfaces.UnmarshalFormatStream(data)
				// 	if err != nil {
				// 		log.Fatal(err)
				// 	}
				// 	flags := []string{
				// 		"--video-title=" + res.Title,
				// 		"--meta-title=" + res.Title,
				// 		"--meta-artist=" + res.Author,
				// 		"--meta-author=" + res.Author,
				// 		"--meta-genre=" + res.Genre,
				// 		"--input-title-format=" + res.Title,
				// 		res.FormatStreams[0].URL,
				// 	}
				// 	if resolution != "" {
				// 		if len(res.AdaptiveFormats) > 1 {
				// 			if _, ok := resolutions[string(resolution)]; !ok {
				// 				fmt.Println("Invalid resolution")
				// 				os.Exit(1)
				// 			}
				// 			for _, v := range res.AdaptiveFormats {
				// 				if string(*v.Container) == string(interfaces.Mp4) && string(*v.Resolution) == resolution {
				// 					flags[len(flags)-1] = v.URL
				// 				}
				// 			}
				// 		}
				// 	}
				// 	if fullscreen {
				// 		println("fullscreen")
				// 		flags = append(flags, "--fullscreen")
				// 	}
				// 	fmt.Println(strings.Join(flags, " "))
				// 	command := exec.Command("vlc", flags...)
				// 	err = command.Start()
				// 	if err != nil {
				// 		log.Fatal(err)
				// 	}
				// 	fmt.Println("vlc opened...")
				// 	os.Exit(0)
				// } else if args[0] == "audio" {
				// 	detailURL := "/api/v1/videos/" + args[1] + "?fields=formatStreams,title,author,genre,adaptiveFormats"
				// 	source, err := instances.FindFastest(&instances.InstanceList, detailURL)
				// 	if err != nil {
				// 		log.Fatal(err)
				// 	}
				// 	// fmt.Println("Requesting " + source.FastestURL)
				// 	defer source.Resp.Body.Close()
				// 	data, err := ioutil.ReadAll(source.Resp.Body)
				// 	res, err := interfaces.UnmarshalFormatStream(data)
				// 	if err != nil {
				// 		log.Fatal(err)
				// 	}
				// 	flags := []string{
				// 		"--video-title=" + res.Title,
				// 		"--meta-title=" + res.Title,
				// 		"--meta-artist=" + res.Author,
				// 		"--meta-author=" + res.Author,
				// 		"--meta-genre=" + res.Genre,
				// 		"--input-title-format=" + res.Title,
				// 	}
				// 	if len(res.AdaptiveFormats) > 1 {
				// 		for _, v := range res.AdaptiveFormats {
				// 			if strings.Contains(v.Type, "audio") && string(*v.Container) == string(interfaces.M4A) {
				// 				flags = append(flags, v.URL)
				// 			}
				// 		}
				// 	} else {
				// 		fmt.Println("Cannot play stream")
				// 		os.Exit(1)
				// 	}
				// 	command := exec.Command("vlc", flags...)
				// 	err = command.Start()
				// 	if err != nil {
				// 		log.Fatal(err)
				// 	}
				// 	fmt.Println("vlc opened...")
				// 	os.Exit(0)
				// }
				fmt.Println("Inside Play")
			}

		} else {
			fmt.Println("No Command")
		}
	},
}

var videoCmd = &cobra.Command{
	Use:   "video [--fullscreen, --resolution] :videoId",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside videoCmd Run with args: %v\n", args)
		detailURL := "/api/v1/videos/" + args[0] + "?fields=formatStreams,title,author,genre,adaptiveFormats,lengthSeconds"
		source, err := instances.FindFastest(&instances.InstanceList, detailURL)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Requesting " + source.FastestURL)
		defer source.Resp.Body.Close()
		data, err := ioutil.ReadAll(source.Resp.Body)
		res, err := interfaces.UnmarshalFormatStream(data)
		if err != nil {
			log.Fatal(err)
		}
		flags := []string{
			"--video-title=" + res.Title,
			"--meta-title=" + res.Title,
			"--meta-artist=" + res.Author,
			"--meta-author=" + res.Author,
			"--meta-genre=" + res.Genre,
			"--input-title-format=" + res.Title,
			res.FormatStreams[0].URL,
		}
		if resolution != "" {
			if len(res.AdaptiveFormats) > 1 {
				if _, ok := resolutions[string(resolution)]; !ok {
					fmt.Println("Invalid resolution")
					os.Exit(1)
				}
				for _, v := range res.AdaptiveFormats {
					if v.Container != nil && *v.Container == interfaces.Webm && v.Resolution != nil && string(*v.Resolution) == resolution && *v.Encoding == "vp9" {
						fmt.Println("v.container")
						flags[len(flags)-1] = v.URL
					}
				}
				flags = append(flags, ":input-slave="+res.AdaptiveFormats[0].URL, ":network-caching=1000")
			}
		}
		if fullscreen {
			println("fullscreen")
			flags = append(flags, "--fullscreen")
		}
		// fmt.Println(strings.Join(flags, " "))
		VLC.Execute(flags...)
	},
}

var audioCmd = &cobra.Command{
	Use:   "audio [no options!] :videoId",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside audioCmd Run with args: %v\n", args)
		detailURL := "/api/v1/videos/" + args[0] + "?fields=formatStreams,title,author,genre,adaptiveFormats,lengthSeconds"
		source, err := instances.FindFastest(&instances.InstanceList, detailURL)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Requesting " + source.FastestURL)
		defer source.Resp.Body.Close()
		data, err := ioutil.ReadAll(source.Resp.Body)
		res, err := interfaces.UnmarshalFormatStream(data)
		if err != nil {
			log.Fatal(err)
		}
		flags := []string{
			"--video-title=" + res.Title,
			"--meta-title=" + res.Title,
			"--meta-artist=" + res.Author,
			"--meta-author=" + res.Author,
			"--meta-genre=" + res.Genre,
			"--input-title-format=" + res.Title,
		}
		if len(res.AdaptiveFormats) > 1 {
			for _, v := range res.AdaptiveFormats {
				if strings.Contains(v.Type, "audio") && string(*v.Container) == string(interfaces.M4A) {
					flags = append(flags, v.URL)
				}
			}
		} else {
			fmt.Println("Cannot play stream")
			os.Exit(1)
		}
		VLC.Execute(flags...)
	},
}

var playlistCmd = &cobra.Command{
	Use:   "playlist [no options!] :playlistId",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside playlistCmd Run with args: %v\n", args)
		playlistURL := "/api/v1/playlists/" + args[0]
		source, err := instances.FindFastest(&instances.InstanceList, playlistURL)
		if err != nil {
			log.Fatal(err)
		}
		defer source.Resp.Body.Close()
		data, err := ioutil.ReadAll(source.Resp.Body)
		res, err := interfaces.UnmarshalPlaylist(data)
		if err != nil {
			log.Fatal(err)
		}
		pl := &vlc.VLCPlaylist{}
		pl.Xmlns = "http://xspf.org/ns/0/"
		pl.Text = "xmlns"
		pl.Vlc = "http://www.videolan.org/vlc/playlist/ns/0/"
		pl.Version = "1"
		pl.Title = "Playlist"
		// pl.Extension.Application = "http://www.videolan.org/vlc/playlist/0"
		Tracks := []vlc.Track{}
		Items := []vlc.ExtensionItem{}
		if len(res.Videos) > 0 {
			fmt.Println("\n Requesting all playlists with " + source.FastestURL + "...")
			playlists := instances.RequestAllPlaylist(source.FastestURL, res.Videos)
			if len(playlists) == 0 {
				fmt.Println("Requested videos not available!")
				os.Exit(1)
			}
			fmt.Println("Total videos: ", len(playlists))
			flags := []string{
				"--network-caching=1000",
				// "--video-title=" + res.Title,
				// "--meta-title=" + res.Title,
				// "--meta-artist=" + res.Author,
				// "--meta-author=" + res.Author,
				// "--input-title-format=" + res.Title,
			}
			for i, v := range playlists {
				id := fmt.Sprint(i)
				localOption := []string{
					"video-title=" + v.Title,
					"input-title-format=" + v.Title,
					"meta-title=" + v.Title,
					"meta-artist=" + v.Author,
					"meta-author=" + v.Author,
				}
				if v != nil {
					if audioOnly {
						if len(v.AdaptiveFormats) > 1 {
							for _, a := range v.AdaptiveFormats {
								if strings.Contains(a.Type, "audio") && string(*a.Container) == string(interfaces.M4A) {
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
				log.Fatal("Cannot create temporary file", err)
			}

			fmt.Println("Created Temporary Playlist File: " + tmpFile.Name())

			// Example writing to the file
			text, err := vlc.MarshalFrom(pl)
			if err != nil {
				log.Fatal("Failed to marshal from data", err)
			}
			if _, err = tmpFile.Write(text); err != nil {
				log.Fatal("Failed to write to temporary file", err)
			}
			flags = append(flags, tmpFile.Name())
			// Remember to clean up the file afterwards
			defer os.Remove(tmpFile.Name())
			VLC.Execute(flags...)
			fmt.Println("Deleting Temporary Playlist")

			// Close the file
			if err := tmpFile.Close(); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("No videos found!")
			os.Exit(1)
		}
	},
}

func init() {
	playCmd.AddCommand(videoCmd, audioCmd, playlistCmd)
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	videoCmd.Flags().BoolVarP(&fullscreen, "fullscreen", "f", false, "Fullscreen video output (default disabled)")
	videoCmd.Flags().StringVarP(&resolution, "resolution", "r", "", "Select high resolution streaming 144p, 240p, 360p, 480p, 720p, 1080p  (default 360p)")
	playlistCmd.Flags().BoolVarP(&audioOnly, "audio-only", "a", false, "Play the playlist in audio format only")
}