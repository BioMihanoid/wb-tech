package command

import "github.com/spf13/cobra"

func InitCmd() *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(initGrepCmd())
	rootCmd.AddCommand(initExit())

	return rootCmd
}
