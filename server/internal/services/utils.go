package services

import (
	"path/filepath"
	"strings"
)

// RemoveFileExtension removes the file extension from a given filename string.
// It returns the filename without the extension.
func RemoveFileExtension(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}
