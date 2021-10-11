package speedtest

import (
	"bytes"
	"text/template"
	"time"
)

type SpeedtestResult struct {
	Type       string      `json:"type"`
	Timestamp  time.Time   `json:"timestamp"`
	Ping       Ping        `json:"ping"`
	Download   SpeedResult `json:"download"`
	Upload     SpeedResult `json:"upload"`
	PacketLoss float64     `json:"packetLoss"`
	Isp        string      `json:"isp"`
	Interface  Interface   `json:"interface"`
	Server     Server      `json:"server"`
	Result     Result      `json:"result"`
}
type Ping struct {
	Jitter  float64 `json:"jitter"`
	Latency float64 `json:"latency"`
}
type SpeedResult struct {
	Bandwidth int `json:"bandwidth"`
	Bytes     int `json:"bytes"`
	Elapsed   int `json:"elapsed"`
}
type Interface struct {
	InternalIP string `json:"internalIp"`
	Name       string `json:"name"`
	MacAddr    string `json:"macAddr"`
	IsVpn      bool   `json:"isVpn"`
	ExternalIP string `json:"externalIp"`
}
type Server struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Country  string `json:"country"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	IP       string `json:"ip"`
}
type Result struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (s *SpeedtestResult) DownloadInMbps() float64 {
	return s.Download.BandwidthInMbps()
}

func (s *SpeedtestResult) UploadInMbps() float64 {
	return s.Upload.BandwidthInMbps()
}

func (s *SpeedResult) BandwidthInMbps() float64 {
	return float64(s.Bandwidth) / float64(125000)
}

func (s *SpeedtestResult) CreateSummary() map[string]interface{} {
	return map[string]interface{}{
		"server_id": s.Server.ID,
		"type":      "network",
		"download": map[string]interface{}{
			"bandwidth": s.DownloadInMbps(),
			"unit":      "mbps",
		},
		"upload": map[string]interface{}{
			"bandwidth": s.UploadInMbps(),
			"unit":      "mbps",
		},
		"ping": map[string]interface{}{
			"latency": s.Ping.Latency,
			"jitter":  s.Ping.Jitter,
		},
		"packet_loss": s.PacketLoss,
		"timestamp":   time.Now(),
		"server": map[string]interface{}{
			"id":       s.Server.ID,
			"name":     s.Server.Name,
			"country":  s.Server.Country,
			"location": s.Server.Location,
			"host":     s.Server.Host,
			"port":     s.Server.Port,
			"ip":       s.Server.IP,
		},
		"image_url": s.Result.URL,
	}
}

func CreateErrorSummary(err error, s Server) map[string]interface{} {
	return map[string]interface{}{
		"server_id": s.ID,
		"type":      "network",
		"download": map[string]interface{}{
			"bandwidth": 0,
			"unit":      "mbps",
		},
		"upload": map[string]interface{}{
			"bandwidth": 0,
			"unit":      "mbps",
		},
		"ping": map[string]interface{}{
			"latency": 0,
			"jitter":  0,
		},
		"packet_loss": 0,
		"timestamp":   time.Now(),
		"error":       err.Error(),
	}
}

func CreateErrorSummaryWithoutServerInfo(err error) map[string]interface{} {
	return map[string]interface{}{
		"server_id": "unknown",
		"download": map[string]interface{}{
			"bandwidth": 0,
			"unit":      "mbps",
		},
		"upload": map[string]interface{}{
			"bandwidth": 0,
			"unit":      "mbps",
		},
		"ping": map[string]interface{}{
			"latency": 0,
			"jitter":  0,
		},
		"packet_loss": 0,
		"timestamp":   time.Now(),
		"error":       err.Error(),
	}
}

type ServerList struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Servers   []Server  `json:"servers"`
}

var t template.Template

func init() {
	t = *template.Must(template.New("test").Parse(`---
{{ .Server.Name }}
- ISP: {{ .Isp }}
- server:      {{ .Server.Name }}
- location:    {{ .Server.Location }}
- download:    {{ .DownloadInMbps }} mbps => megabits/sec ({{ .Download.Bandwidth }} bps => bytes/sec)
- upload:      {{ .UploadInMbps }} mbps => megabits/sec ({{ .Upload.Bandwidth }} bps => bytes/sec)
- ping:        {{ .Ping.Latency }} ms
- jitter:      {{ .Ping.Jitter }} ms
- packet loss: {{ .PacketLoss }} %%
- result:      {{ .Result.URL }}
`))
}

func (r *SpeedtestResult) ToString() string {
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, r); err != nil {
		return err.Error()
	}
	return tpl.String()
}
