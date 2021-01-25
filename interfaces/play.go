package interfaces

import "encoding/json"

func UnmarshalFormatStream(data []byte) (FormatStream, error) {
	var r FormatStream
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FormatStream) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FormatStream struct {
	Title           string                `json:"title"`
	Genre           string                `json:"genre"`
	Author          string                `json:"author"`
	AdaptiveFormats []AdaptiveFormat      `json:"adaptiveFormats"`
	FormatStreams   []FormatStreamElement `json:"formatStreams"`
	LengthSeconds   int64                 `json:"lengthSeconds"`
}

type AdaptiveFormat struct {
	Index          string         `json:"index"`
	Bitrate        string         `json:"bitrate"`
	Init           string         `json:"init"`
	URL            string         `json:"url"`
	Itag           string         `json:"itag"`
	Type           string         `json:"type"`
	Clen           string         `json:"clen"`
	Lmt            string         `json:"lmt"`
	ProjectionType ProjectionType `json:"projectionType"`
	FPS            *int64         `json:"fps,omitempty"`
	Container      *Container     `json:"container,omitempty"`
	Encoding       *string        `json:"encoding,omitempty"`
	Resolution     *Resolution    `json:"resolution,omitempty"`
	QualityLabel   *string        `json:"qualityLabel,omitempty"`
}

type FormatStreamElement struct {
	URL          string `json:"url"`
	Itag         string `json:"itag"`
	Type         string `json:"type"`
	Quality      string `json:"quality"`
	FPS          int64  `json:"fps"`
	Container    string `json:"container"`
	Encoding     string `json:"encoding"`
	Resolution   string `json:"resolution"`
	QualityLabel string `json:"qualityLabel"`
	Size         string `json:"size"`
}

type Container string

const (
	M4A  Container = "m4a"
	Mp4  Container = "mp4"
	Webm Container = "webm"
)

type Resolution string

const (
	R144p  Resolution = "144p"
	R240p  Resolution = "240p"
	R360p  Resolution = "360p"
	R480p  Resolution = "480p"
	R720p  Resolution = "720p"
	R1080p Resolution = "1080p"
)

type ProjectionType string

const (
	Rectangular ProjectionType = "RECTANGULAR"
)
