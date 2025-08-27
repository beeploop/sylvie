package transcoder

type Resolution int

const (
	RES_1080p Resolution = iota
	RES_720p
	RES_480p
	RES_360p
	RES_240p
	RES_144p
)

func (r Resolution) Name() string {
	switch r {
	case 0:
		return "1080p"
	case 1:
		return "720p"
	case 2:
		return "480p"
	case 3:
		return "360p"
	case 4:
		return "240p"
	case 5:
		return "144p"
	default:
		return "unknown resolution"
	}
}

func ResolutionFromName(name string) Resolution {
	switch name {
	case "1080p":
		return RES_1080p
	case "720p":
		return RES_720p
	case "480p":
		return RES_480p
	case "360p":
		return RES_360p
	case "240p":
		return RES_240p
	case "144p":
		return RES_144p
	default:
		return RES_720p
	}
}

type Template struct {
	Resolution   Resolution
	Width        int
	Height       int
	CRF          int
	VideoBitRate int
	AudioBitRate int
	Preset       string
	VideoCodec   string
	AudioCodec   string
}

func TemplateFactory(resolution Resolution) Template {
	switch resolution {
	case RES_1080p:
		return New1080pTemplate()
	case RES_720p:
		return New720pTemplate()
	case RES_480p:
		return New480pTemplate()
	case RES_360p:
		return New360pTemplate()
	case RES_240p:
		return New240pTemplate()
	case RES_144p:
		return New144pTemplate()
	default:
		return New720pTemplate()
	}
}

func New1080pTemplate() Template {
	return Template{
		Resolution:   RES_1080p,
		Width:        1920,
		Height:       1080,
		CRF:          20,
		VideoBitRate: 5_000_000,
		AudioBitRate: 192_000,
		Preset:       "slow",
		VideoCodec:   "libx264",
		AudioCodec:   "aac",
	}
}

func New720pTemplate() Template {
	return Template{
		Resolution:   RES_720p,
		Width:        1280,
		Height:       720,
		CRF:          22,
		VideoBitRate: 2_500_000,
		AudioBitRate: 128_000,
		Preset:       "slow",
		VideoCodec:   "libx264",
		AudioCodec:   "aac",
	}
}

func New480pTemplate() Template {
	return Template{
		Resolution:   RES_480p,
		Width:        854,
		Height:       480,
		CRF:          24,
		VideoBitRate: 1_000_000,
		AudioBitRate: 96_000,
		Preset:       "slow",
		VideoCodec:   "libx264",
		AudioCodec:   "aac",
	}
}

func New360pTemplate() Template {
	return Template{
		Resolution:   RES_360p,
		Width:        640,
		Height:       360,
		CRF:          26,
		VideoBitRate: 800_000,
		AudioBitRate: 64_000,
		Preset:       "slow",
		VideoCodec:   "libx264",
		AudioCodec:   "aac",
	}
}

func New240pTemplate() Template {
	return Template{
		Resolution:   RES_240p,
		Width:        426,
		Height:       240,
		CRF:          28,
		VideoBitRate: 500_000,
		AudioBitRate: 48_000,
		Preset:       "slow",
		VideoCodec:   "libx264",
		AudioCodec:   "aac",
	}
}

func New144pTemplate() Template {
	return Template{
		Resolution:   RES_144p,
		Width:        256,
		Height:       144,
		CRF:          30,
		VideoBitRate: 300_000,
		AudioBitRate: 48_000,
		Preset:       "slow",
		VideoCodec:   "libx264",
		AudioCodec:   "aac",
	}
}
