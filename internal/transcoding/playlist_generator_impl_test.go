package transcoding

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaylistGenerator(t *testing.T) {
	t.Run("test master playlist generator", func(t *testing.T) {
		outDir := t.TempDir()

		input := []Rendetion{
			{
				VideoID:    "1234",
				InputPath:  "path/to/video.mp4",
				Resolution: RES_1080p,
			},
			{
				VideoID:    "1234",
				InputPath:  "path/to/video.mp4",
				Resolution: RES_720p,
			},
			{
				VideoID:    "1234",
				InputPath:  "path/to/video.mp4",
				Resolution: RES_360p,
			},
		}

		expected := []byte("#EXTM3U\n#EXT-X-VERSION:3\n#EXT-X-STREAM-INF:BANDWIDTH=5000000,RESOLUTION=1920x1080\n1080p/index.m3u8\n#EXT-X-STREAM-INF:BANDWIDTH=2800000,RESOLUTION=1280x720\n720p/index.m3u8\n#EXT-X-STREAM-INF:BANDWIDTH=800000,RESOLUTION=640x360\n360p/index.m3u8\n")

		generator := NewPlaylistGenerator(outDir, os.FileMode(0777))

		path, err := generator.Generate(input)
		assert.NoError(t, err)
		assert.NotEmpty(t, path)

		_, err = os.Stat(path)
		assert.NoError(t, err)

		v, err := os.ReadFile(path)
		assert.NoError(t, err, "failed to read generated master playlist")
		assert.Equal(t, len(expected), len(v), "different []byte length between expected and actual")
		assert.EqualValues(t, expected, v)

	})
}
