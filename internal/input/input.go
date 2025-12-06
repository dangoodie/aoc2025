package input

import (
	"os"
	"path/filepath"
)

func Read(relpath string) (string, error) {
	b, err := os.ReadFile(filepath.Clean(relpath))
	if err != nil {
		return "", err
	}

	return string(b), nil
}
