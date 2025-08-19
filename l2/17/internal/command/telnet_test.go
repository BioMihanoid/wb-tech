package command

import (
	"bytes"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func startEchoServer(t *testing.T) (net.Listener, string) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to start echo server: %v", err)
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						return
					}
					_, _ = c.Write(buf[:n])
				}
			}(conn)
		}
	}()

	return ln, ln.Addr().String()
}

func TestTelnet_Echo(t *testing.T) {
	ln, addr := startEchoServer(t)
	defer ln.Close()

	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	rIn, wIn, _ := os.Pipe()
	rOut, wOut, _ := os.Pipe()

	os.Stdin = rIn
	os.Stdout = wOut

	go func() {
		defer wIn.Close()
		_, _ = wIn.Write([]byte("hello\n"))
		time.Sleep(200 * time.Millisecond)
	}()

	flags := &telnetFlags{Timeout: 2 * time.Second}
	host, port, _ := strings.Cut(addr, ":")

	err := runTelnet(flags, host, port)
	if err != nil {
		t.Fatalf("runTelnet failed: %v", err)
	}

	_ = wOut.Close()
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(rOut)

	if !strings.Contains(buf.String(), "hello") {
		t.Errorf("expected echo 'hello', got %q", buf.String())
	}
}

func TestTelnet_ConnectionRefused(t *testing.T) {
	flags := &telnetFlags{Timeout: 500 * time.Millisecond}
	err := runTelnet(flags, "127.0.0.1", "65000")
	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to connect") {
		t.Errorf("unexpected error: %v", err)
	}
}
