package instances

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"MeowTube/client"
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
			res, err := client.Fetch(mirrorURL + p)
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
