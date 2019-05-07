package bunnYmage

import (
	"os"
	"strings"
)

// Check if image exists
func ImageExists(location string) (bool, string, string) {
	var extension string
	var name string
	if stats, err := os.Stat(location); err != nil {
		if os.IsNotExist(err) {
			return false, "", ""
		}
		splitName := strings.Split(stats.Name(), ".")
		extension = "." + splitName[len(splitName)-1]
		name = strings.ReplaceAll(stats.Name(), extension, "")
		extension = strings.ToLower(extension)
	}
	return true, extension, name
}
