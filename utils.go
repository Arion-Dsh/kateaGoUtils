// Package utils ..
package utils

import (
	"path"
	"runtime"
)

// GetCurrentFilePath get file path in current folder
func GetCurrentFilePath(filename string) string {
	_, filePath, _, _ := runtime.Caller(1)
	filePath = path.Join(path.Dir(filePath), filename)
	return filePath
}
