package command

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

type telnetFlags struct {
	Timeout time.Duration
}

func initTelnetCmd() *cobra.Command {
	flags := &telnetFlags{}
	cmd := &cobra.Command{
		Use:   "telnet [host] [port]",
		Short: "Simple telnet client",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			host, port := args[0], args[1]
			return runTelnet(flags, host, port)
		},
	}

	cmd.Flags().DurationVar(&flags.Timeout, "timeout", 10*time.Second, "connection timeout")

	return cmd
}

func runTelnet(flags *telnetFlags, host, port string) error {
	addr := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", addr, flags.Timeout)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	_, _ = fmt.Fprintln(os.Stderr, "Connected to", addr)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		_, _ = io.Copy(os.Stdout, conn)
	}()

	go func() {
		defer wg.Done()
		_, _ = io.Copy(conn, os.Stdin)
	}()

	wg.Wait()
	_, _ = fmt.Fprintln(os.Stderr, "Connection closed")
	return nil
}
