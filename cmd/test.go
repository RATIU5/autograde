package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(testCommand)
}

var testCommand = &cobra.Command{
	Use:   "test",
	Short: "Run existing tests on a project",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
