package transcoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolutions(t *testing.T) {
	t.Run("test resolution name", func(t *testing.T) {
		tests := []struct {
			Name     string
			Input    Resolution
			Expected string
		}{
			{Name: "test 1080p", Input: RES_1080p, Expected: "1080p"},
			{Name: "test 720p", Input: RES_720p, Expected: "720p"},
			{Name: "test 360p", Input: RES_360p, Expected: "360p"},
		}

		for _, tc := range tests {
			name := tc.Input.Name()
			assert.Equal(t, tc.Expected, name)
		}
	})

	t.Run("test ToBitrateSuffixNotation", func(t *testing.T) {
		tests := []struct {
			Input    int
			Expected string
		}{
			{Input: 300_000, Expected: "300k"},
			{Input: 300_000_000, Expected: "300000k"},
			{Input: 1_000, Expected: "1k"},
		}

		for _, tc := range tests {
			v := ToBitrateSuffixNotation(tc.Input)
			assert.Equal(t, tc.Expected, v)
		}
	})
}
