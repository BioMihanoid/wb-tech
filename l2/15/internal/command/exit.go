package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initExitCmd - функция инициализирующая команду exit.
func initExitCmd() *cobra.Command {
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
