package transcoding

import "sylvie/internal/video/entities"

type Probe interface {
	Analyze(inFile string) (entities.VideoMetadata, error)
}
