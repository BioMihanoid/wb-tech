package command

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type flag struct {
	Depth    int
	Parallel int
	Timeout  int
	Output   string
}

// initWgetCmd - функция инициализирующая команду wget.
func initWgetCmd() *cobra.Command {
	flags := &flag{}
	cmd := &cobra.Command{
		Use:   "wget [url]",
		Short: "The wget utility (simplified)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			url := args[0]
			return WgetCmd(flags, url)
		},
	}

	cmd.Flags().IntVarP(&flags.Depth, "depth", "d", 1, "recursion depth")
	cmd.Flags().StringVarP(&flags.Output, "output", "o", "site_mirror", "output directory")
	cmd.Flags().IntVarP(&flags.Parallel, "parallel", "p", 1, "number of parallel downloads")
	cmd.Flags().IntVarP(&flags.Timeout, "timeout", "t", 10, "request timeout in seconds")

	return cmd
}

// WgetCmd — основной обработчик команды
func WgetCmd(flags *flag, url string) error {
	if err := os.MkdirAll(flags.Output, 0755); err != nil {
		return fmt.Errorf("cannot create output dir: %w", err)
	}

	c := NewCrawler(flags.Output, flags.Depth, flags.Parallel, flags.Timeout)

	log.Printf("Start crawling %s (depth=%d, parallel=%d)\n", url, flags.Depth, flags.Parallel)

	if err := c.Crawl(url); err != nil {
		return fmt.Errorf("crawl failed: %w", err)
	}

	log.Println("Done! Site saved to", flags.Output)
	return nil
}
