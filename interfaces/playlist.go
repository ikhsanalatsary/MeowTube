package interfaces

import "encoding/json"

func UnmarshalPlaylist(data []byte) (Playlist, error) {
	var r Playlist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Playlist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Playlist struct {
	Type       string          `json:"type"`
	Title      string          `json:"title"`
	PlaylistID string          `json:"playlistId"`
	Author     string          `json:"author"`
	AuthorID   string          `json:"authorId"`
	Videos     []*VideoElement `json:"videos"`
}
