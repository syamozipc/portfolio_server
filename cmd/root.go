package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "specify mode",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
