package command

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func initEchoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "echo [args...]",
		Short: "Print arguments",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(strings.Join(args, " "))
		},
	}
	return cmd
}
