package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glimcrypt",
	Short: "GlimCrypt is a file encryption and decryption tool",
	Long:  `A tool to encrypt and decrypt files with AES-256 using user-provided keywords.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
}
