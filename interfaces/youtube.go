package interfaces

import (
	"log"
	"net/url"
	"strings"
)

// YoutubeURL map valid url
var YoutubeURL = map[string]string{
	"www.youtube.com": "www.youtube.com",
	"youtu.be":        "youtu.be",
}

// IsValidYoutubeURL tests a string to determine if it is a well-structured url or not.
func IsValidYoutubeURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	// fmt.Println("u", u.Path)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	_, ok := YoutubeURL[u.Host]

	return ok
}

// GetVideoIdFrom youtube url
func GetVideoIdFrom(youtubeURL string) string {
	var videoID string
	u, err := url.Parse(youtubeURL)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	videoID = q.Get("v")
	if videoID == "" {
		videoID = strings.Replace(u.Path, "/", "", 1)
	}

	return videoID
}
