package transcoding

type ProbeStream struct {
	CodecName   string `json:"codec_name"`
	CodecType   string `json:"codec_type"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	FrameRate   string `json:"r_frame_rate"`
	AspectRatio string `json:"display_aspect_ratio"`
}

type ProbeFormat struct {
	Filename   string `json:"filename"`
	Duration   string `json:"duration"`
	Size       string `json:"size"`
	FormatName string `json:"format_name"`
}

type ProbeResult struct {
	Streams []ProbeStream `json:"streams"`
	Format  ProbeFormat   `json:"format"`
}
