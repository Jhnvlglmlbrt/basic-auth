package path

import (
	"path/filepath"
)

func GetFilePath() string {
	absPath, _ := filepath.Abs(".")
	baseDir := filepath.Dir(absPath)
	return baseDir
}
