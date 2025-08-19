package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func initCutCmd() *cobra.Command {
	flag := &flags{}

	cmd := &cobra.Command{
		Use: "cut [options] [STDIN]",
		Short: "The cut command in Linux is designed to extract certain parts of rows (columns) " +
			"from each specified file and send them to the standard output stream.",
		Run: func(cmd *cobra.Command, args []string) {
			cutCmd(flag, os.Stdin, os.Stdout)
		},
	}

	cmd.Flags().StringVarP(&flag.Fields, "fields", "f", "", "specifying the numbers of fields (columns) to be displayed")
	cmd.Flags().StringVarP(&flag.Delimiter, "delimiter", "d", "", "use a different separator (character)")
	cmd.Flags().BoolVarP(&flag.Separated, "separated", "s", false, "only lines containing a separator")

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

func InitCmd() *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(initCutCmd())
	rootCmd.AddCommand(initExit())

	return rootCmd
}
