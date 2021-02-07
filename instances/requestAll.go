package instances

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/corpix/uarand"
	"github.com/ikhsanalatsary/MeowTube/interfaces"
)

// RequestAllPlaylist is a function to find the fastest instance with low latency
func RequestAllPlaylist(url string, videoPlaylists []*interfaces.VideoElement) []*interfaces.FormatStream {
	wg := &sync.WaitGroup{}
	resp := make([]*interfaces.FormatStream, len(videoPlaylists))

	for i, playlist := range videoPlaylists {
		wg.Add(1)
		go func(i int, playlist *interfaces.VideoElement) {
			// fmt.Println("VideoId ", playlist.VideoID)
			jar := NewJar()
			client := &http.Client{Jar: jar}
			req, _ := http.NewRequest("GET", url+"/api/v1/videos/"+playlist.VideoID, nil)
			req.Header.Add("Upgrade-Insecure-Requests", "1")
			req.Header.Add("User-Agent", uarand.GetRandom())
			req.Header.Add("Origin", url)
			// req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
			res, err := client.Do(req)
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
			wg.Done()
		}(i, playlist)
	}
	wg.Wait()

	return resp
}
