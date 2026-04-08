package transcoding

type PlaylistGenerator interface {
	Generate(rendetions []Rendetion) (string, error)
}
