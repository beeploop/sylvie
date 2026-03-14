package transcoding

import (
	"encoding/json"
	"os/exec"
)

type probeImpl struct {
	FfmpegPath string
}

func NewProbeImpl(ffmpegPath string) *probeImpl {
	return &probeImpl{
		FfmpegPath: ffmpegPath,
	}
}

func (p *probeImpl) Analyze(inFile string) (ProbeResult, error) {
	var result ProbeResult

	cmd := p.buildCommand(inFile)

	data, err := cmd.Output()
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (p *probeImpl) buildCommand(inFile string) *exec.Cmd {
	cmd := exec.Command(
		p.FfmpegPath,
		"-v",
		"quiet",
		"-print_format",
		"json",
		"-show_format",
		"-show_streams",
		inFile,
	)

	return cmd
}
