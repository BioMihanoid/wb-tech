package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func initPipeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pipe [cmd1 | cmd2 | ...]",
		Short: "Run commands connected by pipe",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			line := strings.Join(args, " ")
			parts := strings.Split(line, "|")

			var cmds []*exec.Cmd
			for _, part := range parts {
				fields := strings.Fields(strings.TrimSpace(part))
				if len(fields) == 0 {
					continue
				}
				cmds = append(cmds, exec.Command(fields[0], fields[1:]...))
			}

			for i := 0; i < len(cmds)-1; i++ {
				outPipe, err := cmds[i].StdoutPipe()
				if err != nil {
					return fmt.Errorf("pipe error: %w", err)
				}
				cmds[i+1].Stdin = outPipe
			}

			cmds[len(cmds)-1].Stdout = os.Stdout
			cmds[len(cmds)-1].Stderr = os.Stderr

			for _, c := range cmds {
				c.Stderr = os.Stderr
			}

			// запуск всех команд
			for _, c := range cmds {
				if err := c.Start(); err != nil {
					return fmt.Errorf("failed to start: %w", err)
				}
			}
			for _, c := range cmds {
				if err := c.Wait(); err != nil {
					return fmt.Errorf("command failed: %w", err)
				}
			}
			return nil
		},
	}
	return cmd
}
