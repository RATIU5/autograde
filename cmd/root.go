package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Autograde",
	Short: "Autograde can auto-grade code assignments and run tests",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Autograde")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
