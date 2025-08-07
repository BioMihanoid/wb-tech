package command

import (
	"github.com/spf13/cobra"
)

func initSortCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sort",
		Short: "UNIX sort command simple",
		Run:   sortCmd,
	}

	cmd.Flags().IntP("column", "k", -1, "sort by column (starting from 1)")
	cmd.Flags().BoolP("numeric", "n", false, "numeric sort")
	cmd.Flags().BoolP("reverse", "r", false, "reverse sort")
	cmd.Flags().BoolP("unique", "u", false, "unique sort")
	cmd.Flags().StringP("input", "f", "", "input file (defaults to stdin)")
	cmd.Flags().StringP("output", "o", "", "output file (optional)")

	return cmd
}

func InitCmd() *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(initSortCmd())

	return rootCmd
}
