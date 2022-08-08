package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dnCmd = &cobra.Command{
	Use:   "dn",
	Short: "datanode",
	Long:  `Launch datanode for LDFS`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Launching datanode with args...")
	},
}

func init() {
	mainCmd.AddCommand(dnCmd)
}
