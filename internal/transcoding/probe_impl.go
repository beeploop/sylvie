package transcoding

import (
	"encoding/json"
	"os/exec"
	"sylvie/internal/video/entities"
)

type probeImpl struct {
	FfprobePath string
}

func NewProbeImpl(ffmpegPath string) *probeImpl {
	return &probeImpl{
		FfprobePath: ffmpegPath,
	}
}

func (p *probeImpl) Analyze(inFile string) (entities.VideoMetadata, error) {
	cmd := p.buildCommand(inFile)

	data, err := cmd.Output()
	if err != nil {
		return entities.VideoMetadata{}, err
	}

	var result ProbeResult
	if err := json.Unmarshal(data, &result); err != nil {
		return entities.VideoMetadata{}, err
	}

	return normalizeProbeResult(result)
}

func (p *probeImpl) buildCommand(inFile string) *exec.Cmd {
	cmd := exec.Command(
		p.FfprobePath,
		"-v",
		"quiet",
		"-print_format",
		"json",
		"-show_format",
		"-show_streams",
		"-select_streams", "v:0",
		inFile,
	)

	return cmd
}
