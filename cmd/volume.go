package cmd

import (
	"github.com/Raobian/bgofs/pkg/client"
	"github.com/Raobian/bgofs/pkg/common"
	"github.com/Raobian/bgofs/pkg/common/log"
	"github.com/spf13/cobra"
)

var (
	vname string
	vsize string
)

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Volume",
	Long:  `A longer description Volume`,
}

func init() {
	rootCmd.AddCommand(volumeCmd)
	volumeCmd.AddCommand(createCmd)
	CreateParse()
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Volume",
	// Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		vs := common.StrToSize(vsize)
		log.DINFO("create volume name:%s len:%d size:%s size:%d", vname, len(vname), vsize, vs)
		client.CreateVolume(vname, vs)
	},
}

func CreateParse() {
	createCmd.Flags().StringVarP(&vname, "name", "n", "", "Name of volume")
	createCmd.MarkFlagRequired("name")
	createCmd.Flags().StringVarP(&vsize, "size", "s", "", "Size of volume")
	createCmd.MarkFlagRequired("size")
}
