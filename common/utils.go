package common

import (
	"os"
	"path/filepath"
)

func GetBinaryDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "."
	}
	return dir
}
