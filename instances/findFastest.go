package instances

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// FastestInstance is a response the fastest instance from FindFastest
type FastestInstance struct {
	FastestURL string
	Latency    time.Duration
	Resp       http.Response
}

// FindFastest is a function to find the fastest instance with low latency
func FindFastest(urls *[]string, path string) (instance FastestInstance, err error) {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)
	resp := make(chan http.Response)
	resError := make(chan error)

	for _, url := range *urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			res, err := http.Get(mirrorURL + path)
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
				if res.StatusCode >= 200 && res.StatusCode < 400 {
					fmt.Print("Succeed request url: ", mirrorURL+path)
					resp <- *res
					resError <- nil
				} else {
					fmt.Println(mirrorURL + path)
					resError <- errors.New("Unable to request")
				}
			} else {
				resError <- err
			}
			for {
				select {
				case <-resp:
				case <-resError:
					return
				}
			}
		}()
	}

	return FastestInstance{<-urlChan, <-latencyChan, <-resp}, <-resError
}
