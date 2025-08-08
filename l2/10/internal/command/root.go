// Package command - пакет содержащий команды для выполнения в cli.
package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initSortCmd - функция инициализирующая команду sort со всеми нужными параметрами.
func initSortCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sort",
		Short: "UNIX sort command simple",
		Run:   sortCmd,
	}

	cmd.Flags().IntP("column", "k", -1, "sort by column (starting from 1)")
	cmd.Flags().BoolP("numeric", "n", false, "sort by numeric")
	cmd.Flags().BoolP("reverse", "r", false, "reverse sort")
	cmd.Flags().BoolP("unique", "u", false, "unique sort")
	cmd.Flags().StringP("input", "f", "", "input file (defaults to stdin)")
	cmd.Flags().StringP("output", "o", "", "output file (optional)")
	cmd.Flags().BoolP("month", "m", false, "sort by name month")
	cmd.Flags().BoolP("ignore-trailing", "b", false, "ignore trailing blanks")
	cmd.Flags().BoolP("check", "c", false, "check whether input is sorted")
	cmd.Flags().BoolP("human-numeric", "H", false, "human-readable numeric sort "+
		"(e.g. 2K, 1M)")

	return cmd
}

// initExit - функция инициализирующая команду exit.
func initExit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exit",
		Short: "exit app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("process finished")
			os.Exit(0)
		},
	}

	return cmd
}

// InitCmd - инициализация всех команд.
func InitCmd() *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(initSortCmd())
	rootCmd.AddCommand(initExit())

	return rootCmd
}
