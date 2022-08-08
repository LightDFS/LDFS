package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nnCmd = &cobra.Command{
	Use:   "nn",
	Short: "namenode",
	Long:  `Launch namenode for LDFS`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Launching namenode with args...")
	},
}

func init() {
	mainCmd.AddCommand(nnCmd)
}
