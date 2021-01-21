package instances

// InstanceList is list of Invidious instance sites
var InstanceList = []string{
	// "https://invidious.snopyta.org",
	// "https://yewtu.be",
	"https://invidious.tube", "https://invidious.xyz",
	"https://invidious.kavin.rocks",
	// "https://tube.connect.cafe",
	"https://invidious.zapashcanon.fr",
	// "https://invidious.fdn.fr",
	"https://invidiou.site",
	// "https://vid.mint.lgbt",
	"https://invidious.site", "https://invidious.048596.xyz",
}

// // FastestInstance is a response the fastest instance from FindFastest
// type FastestInstance struct {
// 	FastestURL string        `json:"fastest_url"`
// 	Latency    time.Duration `json:"latency"`
// }

// // FindFastest is a function to find the fastest instance with low latency
// func FindFastest(urls *[]string) FastestInstance {
// 	urlChan := make(chan string)
// 	latencyChan := make(chan time.Duration)

// 	for _, url := range *urls {
// 		mirrorURL := url
// 		go func() {
// 			start := time.Now()
// 			_, err := http.Get(mirrorURL + "/feed/popular")
// 			latency := time.Now().Sub(start) / time.Millisecond
// 			if err == nil {
// 				urlChan <- mirrorURL
// 				latencyChan <- latency
// 			}
// 		}()
// 	}
// 	return FastestInstance{<-urlChan, <-latencyChan}
// }
