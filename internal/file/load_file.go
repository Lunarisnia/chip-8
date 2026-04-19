package file

import (
	"os"
	"path"
	"path/filepath"
)

func ReadFile(filePath string) ([]byte, error) {
	if !filepath.IsAbs(filePath) {
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		filePath = path.Join(dir, filePath)
	}
	bin, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return bin, nil
}
