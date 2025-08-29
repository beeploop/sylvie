package rabbitmq

type Message struct {
	VideoID     string   `json:"video_id"`
	Path        string   `json:"path"`
	Resolutions []string `json:"resolutions"`
}
