package transcoding

import (
	"sylvie/internal/video/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProbe(t *testing.T) {
	t.Run("test probe result normalization", func(t *testing.T) {
		tests := []struct {
			Name     string
			Input    ProbeResult
			Expected entities.VideoMetadata
		}{
			{
				Name: "test happy path",
				Input: ProbeResult{
					Streams: []ProbeStream{
						{
							CodecName:   "h264",
							CodecType:   "video",
							Width:       1920,
							Height:      1080,
							FrameRate:   "60/1",
							AspectRatio: "16:9",
						},
					},
					Format: ProbeFormat{
						Filename:   "test.mp4",
						Duration:   "10.0000",
						FormatName: "mov,mp4",
					},
				},
				Expected: entities.VideoMetadata{
					Width:     1920,
					Height:    1080,
					Framerate: 60,
					Duration:  10,
					Codec:     "h264",
				},
			},
		}

		for _, tc := range tests {
			t.Run(tc.Name, func(t *testing.T) {
				result, err := normalizeProbeResult(tc.Input)
				assert.NoError(t, err)
				assert.EqualValues(t, tc.Expected, result)
			})
		}
	})
}
