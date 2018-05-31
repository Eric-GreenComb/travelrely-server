package tools

import (
	"path"
	"strings"
)

// GetFileName GetFileName
func GetFileName(file string) (filename string, ext string) {
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(file)
	ext = path.Ext(filenameWithSuffix)
	filename = strings.TrimSuffix(filenameWithSuffix, ext)
	return
}
