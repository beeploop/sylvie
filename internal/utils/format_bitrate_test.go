package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatBitrate(t *testing.T) {
	t.Run("test format bitrate return", func(t *testing.T) {
		result1 := BitrateToFfmpegStyleString(300_000)
		require.Equal(t, "300k", result1)

		result2 := BitrateToFfmpegStyleString(2_500_000)
		require.Equal(t, "2500k", result2)
	})
}
