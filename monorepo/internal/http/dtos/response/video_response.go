package response

type VideoResponse struct {
	VideoID  string `json:"video_id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Status   string `json:"status"`
}
