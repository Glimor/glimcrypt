package cmd

import (
	"fmt"
	"glimmer/internal"
	"log"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts a file",
	Run: func(cmd *cobra.Command, args []string) {
		keywords := internal.GetUserKeywords()
		key, err := internal.GenerateKeyFromKeywords(keywords)
		if err != nil {
			log.Fatalf("Error generating key: %v", err)
		}

		filename := internal.SelectFile()

		err = internal.EncryptFile(filename, key)
		if err != nil {
			log.Fatalf("Error encrypting file: %v", err)
		}
		fmt.Println("File encrypted successfully.")
	},
}
