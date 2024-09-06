package main

import (
	"glimmer/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error starting app: %v", err)
	}
}
