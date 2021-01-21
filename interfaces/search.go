package interfaces

import (
	"encoding/json"

	json2 "github.com/nwidger/jsoncolor"
)

type Search []SearchElement

func UnmarshalSearch(data []byte) (Search, error) {
	var r Search
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Search) Marshal() ([]byte, error) {
	return json2.MarshalIndent(r, "", " ")
}

type SearchElement struct {
	Type          Type    `json:"type"`
	Title         string  `json:"title"`
	VideoID       string  `json:"videoId"`
	Author        string  `json:"author"`
	AuthorID      string  `json:"authorId"`
	PlaylistID    *string `json:"playlistId,omitempty"`
	PublishedText string  `json:"publishedText"`
}

type Type string

const (
	ChannelType  Type = "channel"
	PlaylistType Type = "playlist"
	VideoType    Type = "video"
)
