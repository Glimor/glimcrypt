package internal

import (
	"log"

	"github.com/ncruces/zenity"
)

func SelectFile() string {
	filename, err := zenity.SelectFile()
	if err != nil {
		log.Fatalf("Failed to select file: %v", err)
	}
	return filename
}
