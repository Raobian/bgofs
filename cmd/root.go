package cmd

import (
	"os"

	"github.com/Raobian/bgofs/pkg/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bgofs",
	Short: "A brief description",
	Long:  `A longer description`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(server.Init)
}
