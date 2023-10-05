package cmd

import (
	"fmt"

	"github.com/RATIU5/autograde/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of autograde",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("autograde v%v\n", config.GetVersion())
	},
}
