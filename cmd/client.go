package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client",
	Long:  `Launch client for LDFS`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Launching client with args...")
	},
}

func init() {
	mainCmd.AddCommand(clientCmd)
}
