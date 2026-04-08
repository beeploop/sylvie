package transcoding

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type playlistGeneratorImpl struct {
	BaseDir    string
	Permission os.FileMode
}

func NewPlaylistGenerator(baseDir string, permission os.FileMode) *playlistGeneratorImpl {
	return &playlistGeneratorImpl{
		BaseDir:    baseDir,
		Permission: permission,
	}
}

func (g *playlistGeneratorImpl) Generate(rendetions []Rendetion) (string, error) {
	outDir := filepath.Join(g.BaseDir, rendetions[0].VideoID)
	if err := os.MkdirAll(outDir, g.Permission); err != nil {
		return "", err
	}

	var b strings.Builder

	b.WriteString("#EXTM3U\n")
	b.WriteString("#EXT-X-VERSION:3\n")

	for _, rendetion := range rendetions {
		b.WriteString(
			fmt.Sprintf(
				"#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n",
				rendetion.Resolution.VideoBitRate(),
				rendetion.Resolution.Dimension()),
		)

		b.WriteString(fmt.Sprintf("%s/index.m3u8\n", rendetion.Resolution.Name()))
	}

	path := filepath.Join(outDir, "master.m3u8")
	if err := os.WriteFile(path, []byte(b.String()), g.Permission); err != nil {
		return "", err
	}

	return path, nil
}
