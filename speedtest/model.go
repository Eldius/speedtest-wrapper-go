package speedtest

import "time"

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
		"download":      s.DownloadInMbps(),
		"download_unit": "mbps",
		"upload":        s.UploadInMbps(),
		"upload_unit":   "mbps",
		"ping":          s.Ping.Latency,
		"jitter":        s.Ping.Jitter,
		"packet_loss":   s.PacketLoss,
		"timestamp":     time.Now(),
	}
}
