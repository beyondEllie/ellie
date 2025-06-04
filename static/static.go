package static

import (
	"embed"
	"io/fs"
)

// go:embed *
var Static embed.FS

func GetStaticFile(path string) (string, error) {
	content, err := fs.ReadFile(Static, path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}