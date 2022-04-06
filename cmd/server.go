package cmd

import (
	"github.com/Raobian/bgofs/pkg/server"

	"github.com/spf13/cobra"
)

// fuseCmd represents the fuse command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server",
	Long:  `A longer description Server`,
	Run: func(cmd *cobra.Command, args []string) {
		// mp, _ := cmd.Flags().GetString("mountpoint")
		server.Server()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	// serverCmd.PersistentFlags().StringP("mountpoint", "m", "/mnt/bane_fuse", "Mountpoint")
}
