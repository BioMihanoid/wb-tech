package command

import (
	"fmt"
	"os"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

func initKillCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kill <pid>",
		Short: "Terminate process by PID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pid, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid PID: %w", err)
			}
			process, err := os.FindProcess(pid)
			if err != nil {
				return fmt.Errorf("cannot find process: %w", err)
			}
			if err := process.Signal(syscall.SIGKILL); err != nil {
				return fmt.Errorf("failed to kill process: %w", err)
			}
			fmt.Printf("Process %d killed\n", pid)
			return nil
		},
	}

	return cmd
}
