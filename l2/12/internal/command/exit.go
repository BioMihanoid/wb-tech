package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
