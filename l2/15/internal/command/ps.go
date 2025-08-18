package command

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func initPsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ps",
		Short: "List running processes",
		RunE: func(cmd *cobra.Command, args []string) error {
			out, err := exec.Command("ps", "aux").Output()
			if err != nil {
				return fmt.Errorf("failed to run ps: %w", err)
			}
			fmt.Print(string(out))
			return nil
		},
	}

	return cmd
}
