package transcoding

type ThumbnailInput struct {
	VideoID  string
	Filepath string
}

type ThumbnailGenerator interface {
	Generate(ThumbnailInput) (string, error)
}
