package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func initRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run [command]...",
		Short: "Run external command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c := exec.Command(args[0], args[1:]...)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Stdin = os.Stdin

			if err := c.Run(); err != nil {
				return fmt.Errorf("command failed: %w", err)
			}
			return nil
		},
	}

	return cmd
}
