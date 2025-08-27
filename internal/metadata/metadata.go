package metadata

type Rendetion struct {
	Resolution string  `json:"resolution"`
	Width      int     `json:"width"`
	Height     int     `json:"height"`
	FrameRate  float32 `json:"frame_rate"`
	Bitrate    int     `json:"bitrate"`
	Codec      string  `json:"codec"`
	FileSize   int     `json:"file_size"`
	Path       string  `json:"path"`
}

type SourceMetadata struct {
	Duration         float32 `json:"duration"`
	AspectRatio      string  `json:"aspect_ratio"`
	VideoCodec       string  `json:"video_codec"`
	AudioCodec       string  `json:"audio_codec"`
	FrameRate        float32 `json:"frame_rate"`
	OriginalFileSize int     `json:"original_file_size"`
}

type Thumbnail struct {
	Timestamp int    `json:"timestamp"`
	Type      string `json:"type"` // default | spritesheet
	Path      string `json:"path"`
}

type Metadata struct {
	VideoID        string         `json:"video_id"`
	SourceMetadata SourceMetadata `json:"source_metadata"`
	Rendetions     []Rendetion    `json:"rendetions"`
	Thumbnails     []Thumbnail    `json:"thumbnails"`
}
