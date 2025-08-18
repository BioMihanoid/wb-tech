package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func initPwdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pwd",
		Short: "Print working directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			fmt.Println(dir)
			return nil
		},
	}
	return cmd
}
