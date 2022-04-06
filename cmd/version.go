package cmd

import (
	"fmt"
	"os"

	"github.com/Raobian/bgofs/pkg/version"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		// output, err := ExecuteCommand("git", "rev-parse", "--short", "HEAD")
		// if err != nil {
		// 	Error(cmd, args, err)
		// }

		fmt.Fprint(os.Stdout, version.Version()+"\n")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
