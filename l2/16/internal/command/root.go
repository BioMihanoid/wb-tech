package command

import "github.com/spf13/cobra"

// InitCmd - инициализация всех команд приложения.
func InitCmd() *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(initWgetCmd())
	rootCmd.AddCommand(initExitCmd())

	return rootCmd
}
