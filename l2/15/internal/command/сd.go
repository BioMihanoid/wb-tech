package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func initCdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cd [dir]",
		Short: "Change current directory",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := os.Chdir(args[0]); err != nil {
				return fmt.Errorf("cd failed: %w", err)
			}
			return nil
		},
	}
	return cmd
}
