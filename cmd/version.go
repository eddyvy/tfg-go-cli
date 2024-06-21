package cmd

import (
	"fmt"

	"github.com/eddyvy/tfg-go-cli/internal"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the installed version of tfg",
	Run:   runVersion(fmt.Println),
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(print func(a ...interface{}) (n int, err error)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		print("Version:", internal.TFG_VERSION)
	}
}
