package transcoder

type Transcoder interface {
	Transcode(Rendetion) (string, error)
}
