package config

type App struct {
	FfmpegPath  string `yaml:"ffmpeg_path"`
	FfprobePath string `yaml:"ffprobe_path"`
	TempDir     string `yaml:"temp_dir"`
}

type ImageVariants struct {
	Name   string `yaml:"name"`
	Width  int    `yaml:"width"`
	Format string `yaml:"format"`
}

type ImageProcessing struct {
	Variants []ImageVariants `yaml:"variants"`
}

type VideoVariants struct {
	Name         string `yaml:"name"`
	Width        int    `yaml:"width"`
	Height       int    `yaml:"height"`
	VideoBitRate string `yaml:"video_bitrate"`
	AudioBitRate string `yaml:"audio_bitrate"`
	Format       string `yaml:"format"`
}

type VideoThumbnail struct {
	TimeSeconds int    `yaml:"time_seconds"`
	Width       int    `yaml:"width"`
	Format      string `yaml:"format"`
}

type VideoProcessing struct {
	Variants  []VideoVariants `yaml:"variants"`
	Thumbnail VideoThumbnail  `yaml:"thumbnail"`
}

type Processing struct {
	Image ImageProcessing `yaml:"image"`
	Video VideoProcessing `yaml:"video"`
}

type Config struct {
	App        App        `yaml:"app"`
	Processing Processing `yaml:"processing"`
}
