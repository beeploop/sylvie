package queue

type Job struct {
	VideoID string `json:"video_id"`
	Path    string `json:"path"`
}
