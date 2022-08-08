package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var mainCmd = &cobra.Command{
	Use:   "ldfs",
	Short: "CLI",
	Long:  `CLI launcher for LDFS namenode, datanode, and client`,
}

func Execute() {
	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
