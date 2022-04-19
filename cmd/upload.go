package cmd

import (
	"github.com/Raobian/bgofs/pkg/client"

	"github.com/spf13/cobra"
)

var fname string

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload",
	Long:  `A longer description Upload`,
	Run: func(cmd *cobra.Command, args []string) {
		client.Upload(fname)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&fname, "filename", "f", "", "File name to upload, (required)")
	uploadCmd.MarkFlagRequired("filename")
}
