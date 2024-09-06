package cmd

import (
	"fmt"
	"glimmer/internal"
	"log"

	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts a file",
	Run: func(cmd *cobra.Command, args []string) {

		keywords := internal.GetUserKeywords()
		key, err := internal.GenerateKeyFromKeywords(keywords)
		if err != nil {
			log.Fatalf("Error generating key: %v", err)
		}

		filename := internal.SelectFile()

		err = internal.DecryptFile(filename, key)
		if err != nil {
			log.Fatalf("Invalid 8 keyword provided or file is not encrypted with GlimCrypt.")
		}
		fmt.Println("File decrypted successfully.")
	},
}
