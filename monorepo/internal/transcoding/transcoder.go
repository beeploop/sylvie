package transcoding

type Transcoder interface {
	Transcode(Rendetion) (string, error)
}
