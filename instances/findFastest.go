package instances

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/corpix/uarand"
	"github.com/ikhsanalatsary/MeowTube/client"
)

// FastestInstance is a response the fastest instance from FindFastest
type FastestInstance struct {
	FastestURL string
	Latency    time.Duration
	Resp       *http.Response
	Error      error
}

// FindFastest is a function to find the fastest instance with low latency
func FindFastest(path string) FastestInstance {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)
	resp := make(chan *http.Response)
	resError := make(chan error)
	instanceUrls, err := FindInstanceList()

	if err != nil {
		return FastestInstance{
			Error: err,
		}
	}

	for _, url := range instanceUrls {
		mirrorURL := url
		go func() {
			start := time.Now()
			// There is an videoID that uses `-` in the leading of their characters.
			// But go cannot remove single quotes on string characters automatically.
			p := strings.Replace(path, "'", "", 2)
			req, _ := http.NewRequest("GET", mirrorURL+p, nil)
			req.Header.Add("Upgrade-Insecure-Requests", "1")
			req.Header.Add("User-Agent", uarand.GetRandom())
			req.Header.Add("Origin", mirrorURL)
			req.Header.Add("Accept", "*/*")
			// req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
			res, err := client.Request.Do(req)
			latency := time.Now().Sub(start) / time.Millisecond
			urlChan <- mirrorURL
			latencyChan <- latency
			if err == nil {
				if res.StatusCode >= 200 && res.StatusCode < 400 {
					fmt.Print("Succeed request url: ", mirrorURL+p)
					resp <- res
					resError <- nil
				} else {
					fmt.Println("Failed request url", mirrorURL+p)
					fmt.Println("statusCode: ", res.StatusCode)
					resp <- nil
					resError <- errors.New("Unable to request")
				}
			} else {
				fmt.Println("Failed request url", mirrorURL+p)
				resp <- nil
				resError <- err
			}
		}()
	}

	return FastestInstance{<-urlChan, <-latencyChan, <-resp, <-resError}
}
