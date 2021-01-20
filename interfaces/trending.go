package interfaces

import (
	"encoding/json"

	json2 "github.com/nwidger/jsoncolor"
)

type Trending []TrendingElement

func UnmarshalTrending(data []byte) (Trending, error) {
	var r Trending
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Trending) Marshal() ([]byte, error) {
	return json2.MarshalIndent(r, "", " ")
}

type TrendingElement struct {
	Type      Type   `json:"type"`
	Title     string `json:"title"`
	VideoID   string `json:"videoId"`
	Author    string `json:"author"`
	AuthorID  string `json:"authorId"`
	AuthorURL string `json:"authorUrl"`
	// VideoThumbnails []VideoThumbnail `json:"videoThumbnails"`
	// Description     string           `json:"description"`
	// DescriptionHTML string `json:"descriptionHtml"`
	ViewCount     int64  `json:"viewCount"`
	Published     int64  `json:"published"`
	PublishedText string `json:"publishedText"`
	LengthSeconds int64  `json:"lengthSeconds"`
	LiveNow       bool   `json:"liveNow"`
	Paid          bool   `json:"paid"`
	Premium       bool   `json:"premium"`
	IsUpcoming    bool   `json:"isUpcoming"`
}

type VideoThumbnail struct {
	Quality Quality `json:"quality"`
	URL     string  `json:"url"`
	Width   int64   `json:"width"`
	Height  int64   `json:"height"`
}

type Quality string

const (
	Default       Quality = "default"
	End           Quality = "end"
	High          Quality = "high"
	Maxres        Quality = "maxres"
	Maxresdefault Quality = "maxresdefault"
	Medium        Quality = "medium"
	Middle        Quality = "middle"
	Sddefault     Quality = "sddefault"
	Start         Quality = "start"
)
