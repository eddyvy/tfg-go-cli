package cmd

import (
	"fmt"

	"github.com/eddyvy/tfg-go-cli/internal"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the installed version of tfg",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", internal.TFG_VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
