package transcoding

type Probe interface {
	Analyze(inFile string) (ProbeResult, error)
}
