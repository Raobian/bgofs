package cmd

import (
	"github.com/Raobian/bgofs/pkg/client"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "Upload",
	Short: "Upload",
	Long:  `A longer description Upload`,
	Run: func(cmd *cobra.Command, args []string) {
		fname, err := cmd.Flags().GetString("filename")
		if err != nil {
			Error(cmd, args, err)
		}
		client.Run(fname)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.PersistentFlags().StringP("filename", "m", "", "File name to upload")
}
