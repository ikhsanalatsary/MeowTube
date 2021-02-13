package instances

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// InstanceList is list of Invidious instance sites, currently we used from uptime
var InstanceList = []string{
	"https://invidious.snopyta.org",
	"https://yewtu.be",
	"https://invidious.tube", "https://invidious.xyz",
	"https://invidious.kavin.rocks",
	"https://tube.connect.cafe",
	"https://invidious.zapashcanon.fr",
	"https://invidious.fdn.fr",
	"https://invidiou.site",
	"https://vid.mint.lgbt",
	"https://invidious.site",
	"https://invidious.048596.xyz",
}

// UnmarshalServerInstanceList unmarshaling response data to ServerInstanceList
func UnmarshalServerInstanceList(data []byte) (ServerInstanceList, error) {
	var r ServerInstanceList
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal marshaling ServerInstanceList to response data
func (r *ServerInstanceList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ServerInstanceList uptime response
type ServerInstanceList struct {
	Status     string     `json:"status"`
	Psp        Psp        `json:"psp"`
	Days       []string   `json:"days"`
	Statistics Statistics `json:"statistics"`
}

// Psp ServerInstanceList uptime response ServerInstanceList.Psp
type Psp struct {
	PerPage       int64         `json:"perPage"`
	TotalMonitors int64         `json:"totalMonitors"`
	Monitors      []Monitor     `json:"monitors"`
	Timezone      string        `json:"timezone"`
	Logs          []interface{} `json:"logs"`
}

// Monitor ServerInstanceList uptime response ServerInstanceList.Monitors
type Monitor struct {
	MonitorID   int64       `json:"monitorId"`
	CreatedAt   int64       `json:"createdAt"`
	StatusClass StatusClass `json:"statusClass"`
	Name        string      `json:"name"`
	URL         interface{} `json:"url"`
	Type        Type        `json:"type"`
	DailyRatios []L1        `json:"dailyRatios"`
	The90DRatio L1          `json:"90dRatio"`
	The30DRatio L1          `json:"30dRatio"`
}

// L1 ServerInstanceList uptime response ServerInstanceList.Monitors.DailyRatios
type L1 struct {
	Ratio string `json:"ratio"`
	Label Label  `json:"label"`
}

// Statistics ServerInstanceList uptime response ServerInstanceList.Statistics
type Statistics struct {
	Uptime         Uptime      `json:"uptime"`
	LatestDowntime interface{} `json:"latest_downtime"`
	Counts         Counts      `json:"counts"`
	CountResult    string      `json:"count_result"`
}

// Counts ServerInstanceList uptime response ServerInstanceList.Statistics.Counts
type Counts struct {
	Up     int64 `json:"up"`
	Down   int64 `json:"down"`
	Paused int64 `json:"paused"`
}

// Uptime ServerInstanceList uptime response ServerInstanceList.Statistics.Uptime
type Uptime struct {
	L1  L1 `json:"l1"`
	L7  L1 `json:"l7"`
	L30 L1 `json:"l30"`
	L90 L1 `json:"l90"`
}

type Label string

const (
	LabelSuccess Label = "success"
	LabelWarning Label = "warning"
	LabelDanger  Label = "danger"
)

type StatusClass string

const (
	StatusDanger  StatusClass = "danger"
	StatusSuccess StatusClass = "success"
	StatusWarning StatusClass = "warning"
)

type Type string

const (
	HTTPS Type = "HTTP(s)"
)

// exludeNames are not the list of invidious instances
var excludeNames = map[string]string{
	"api.invidious.io": "api.invidious.io",
	"invidious.io":     "invidious.io",
	"invidio.us":       "invidio.us",
}

// FindInstanceList to find Instance list from uptime monitor invidious server
func FindInstanceList() (urls []string, err error) {
	var data []byte
	var resp ServerInstanceList
	var instanceURLs []string
	url := "https://uptime.invidious.io/api/getMonitorList/89VnzSKAn?page=1&_=1611588676444"
	res, err := http.Get(url)
	if err == nil {
		if res.StatusCode >= 200 && res.StatusCode < 400 {
			defer res.Body.Close()
			data, err = ioutil.ReadAll(res.Body)
			resp, err = UnmarshalServerInstanceList(data)
			if resp.Status == "ok" {
				for _, v := range resp.Psp.Monitors {
					if v.StatusClass == StatusSuccess {
						if _, exist := excludeNames[v.Name]; !exist {
							if n := viper.Get(v.Name); n != nil && n == false {
								continue
							}
							instanceURLs = append(instanceURLs, "https://"+v.Name)
						}
					}
				}
			} else {
				err = errors.New("Server not ok")
			}
		} else {
			err = errors.New("Server down")
		}
	}
	return instanceURLs, err
}
