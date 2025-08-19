package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"wb-tech/l2/15/internal/command"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")

	for {
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		rootCmd := command.InitCmd()
		rootCmd.SetArgs(parts)

		if err := rootCmd.Execute(); err != nil {
			fmt.Println(err)
		}
		fmt.Print("> ")
	}
}
