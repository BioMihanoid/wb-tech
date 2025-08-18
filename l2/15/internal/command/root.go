package command

import "github.com/spf13/cobra"

// InitCmd - инициализация всех команд приложения.
func InitCmd() *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(initExitCmd())
	rootCmd.AddCommand(initCdCmd())
	rootCmd.AddCommand(initPwdCmd())
	rootCmd.AddCommand(initEchoCmd())
	rootCmd.AddCommand(initKillCmd())
	rootCmd.AddCommand(initPsCmd())
	rootCmd.AddCommand(initRunCmd())
	rootCmd.AddCommand(initPipeCmd())

	return rootCmd
}
