package instances

import (
	"io/ioutil"
	"sync"

	"MeowTube/client"
	"MeowTube/interfaces"
)

// RequestAllPlaylist is a function to find the fastest instance with low latency
func RequestAllPlaylist(url string, videoPlaylists []*interfaces.VideoElement) []*interfaces.FormatStream {
	wg := &sync.WaitGroup{}
	resp := make([]*interfaces.FormatStream, len(videoPlaylists))

	for i, playlist := range videoPlaylists {
		wg.Add(1)
		go func(i int, playlist *interfaces.VideoElement) {
			// fmt.Println("VideoId ", playlist.VideoID)
			res, err := client.Fetch(url + "/api/v1/videos/" + playlist.VideoID)
			defer wg.Done()
			if err == nil {
				if res.StatusCode >= 200 && res.StatusCode < 400 {
					defer res.Body.Close()
					data, err := ioutil.ReadAll(res.Body)
					result, err := interfaces.UnmarshalFormatStream(data)
					if err == nil {
						resp[i] = &result
					}
				}
			}
		}(i, playlist)
	}
	wg.Wait()

	return resp
}
