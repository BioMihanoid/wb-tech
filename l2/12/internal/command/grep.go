package command

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

type grepFlags struct {
	After       int
	Before      int
	Context     int
	IgnoreCase  bool
	Count       bool
	InvertMatch bool
	Fixed       bool
	LineNumber  bool
}

func initGrepCmd() *cobra.Command {
	flags := &grepFlags{}

	cmd := &cobra.Command{
		Use:   "grep [flags] PATTERN [FILE...]",
		Short: "Search for PATTERN in each FILE or standard input",
		Args:  cobra.MinimumNArgs(1),
		PreRun: func(_ *cobra.Command, _ []string) {
			normalizeContextFlags(flags)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			pattern := args[0]
			files := args[1:]
			return grepCmd(flags, pattern, files, os.Stdin, os.Stdout)
		},
	}

	addGrepFlags(cmd, flags)
	return cmd
}

func addGrepFlags(cmd *cobra.Command, flags *grepFlags) {
	cmd.Flags().IntVarP(&flags.After, "after", "A", 0, "print N lines of trailing context after matching lines")
	cmd.Flags().IntVarP(&flags.Before, "before", "B", 0, "print N lines of leading context before matching lines")
	cmd.Flags().IntVarP(&flags.Context, "context", "C", 0, "print N lines of output context (both before and after)")

	cmd.Flags().BoolVarP(&flags.Count, "count", "c", false, "print only a count of matching lines")
	cmd.Flags().BoolVarP(&flags.IgnoreCase, "ignore-case", "i", false, "ignore case distinctions")
	cmd.Flags().BoolVarP(&flags.InvertMatch, "invert-match", "v", false, "select non-matching lines")
	cmd.Flags().BoolVarP(&flags.Fixed, "fixed-strings", "F", false, "interpret PATTERN as a fixed string")
	cmd.Flags().BoolVarP(&flags.LineNumber, "line-number", "n", false, "print line number with output lines")
}

func normalizeContextFlags(flags *grepFlags) {
	if flags.Context > 0 {
		if flags.After == 0 {
			flags.After = flags.Context
		}
		if flags.Before == 0 {
			flags.Before = flags.Context
		}
	}
}

// grepCmd — точка входа: читает из файлов или STDIN и печатает результат.
func grepCmd(flags *grepFlags, pattern string, files []string, stdin io.Reader, stdout io.Writer) error {
	matchFunc, err := buildMatcher(pattern, flags.Fixed, flags.IgnoreCase)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "invalid pattern:", err)
		return nil
	}

	if len(files) == 0 {
		lines, err := readAllLines(stdin)
		if err != nil {
			return err
		}
		printGrep(flags, lines, matchFunc, "", stdout)
		return nil
	}

	for _, name := range files {
		f, err := os.Open(name)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cannot open %s: %v\n", name, err)
			continue
		}
		lines, err := readAllLines(f)
		_ = f.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "read error %s: %v\n", name, err)
			continue
		}
		printGrep(flags, lines, matchFunc, name, stdout)
	}
	return nil
}

// buildMatcher — конструирует функцию сопоставления строки с шаблоном.
func buildMatcher(pattern string, fixed, ignoreCase bool) (func(string) bool, error) {
	if fixed {
		if ignoreCase {
			p := strings.ToLower(pattern)
			return func(s string) bool {
				return strings.Contains(strings.ToLower(s), p)
			}, nil
		}
		return func(s string) bool {
			return strings.Contains(s, pattern)
		}, nil
	}
	if ignoreCase {
		pattern = "(?i)" + pattern
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.MatchString, nil
}

func printGrep(flags *grepFlags, lines []string, match func(string) bool, filename string, w io.Writer) {
	matchIdx, contextIdx := computeMatchesAndContext(flags, lines, match)

	if flags.Count {
		count := len(matchIdx)
		if filename != "" {
			_, _ = fmt.Fprintf(w, "%s:%d\n", filename, count)
		} else {
			_, _ = fmt.Fprintln(w, count)
		}
		return
	}

	toPrint := make(map[int]struct{}, len(matchIdx)+len(contextIdx))
	for i := range matchIdx {
		toPrint[i] = struct{}{}
	}
	for i := range contextIdx {
		toPrint[i] = struct{}{}
	}
	indexes := make([]int, 0, len(toPrint))
	for i := range toPrint {
		indexes = append(indexes, i)
	}
	sort.Ints(indexes)

	for _, i := range indexes {
		if filename != "" {
			_, _ = fmt.Fprint(w, filename, ":")
		}
		if flags.LineNumber {
			_, _ = fmt.Fprintf(w, "%d:", i+1)
		}
		_, _ = fmt.Fprintln(w, lines[i])
	}
}

// computeMatchesAndContext — возвращает:
//  1. map[int]struct{} множество индексов совпадений,
//  2. map[int]struct{} множество индексов контекста (до/после)
func computeMatchesAndContext(flags *grepFlags, lines []string, match func(string) bool) (map[int]struct{}, map[int]struct{}) {
	matchIdx := make(map[int]struct{})
	contextIdx := make(map[int]struct{})

	for i, line := range lines {
		ok := match(line)
		if flags.InvertMatch {
			ok = !ok
		}
		if !ok {
			continue
		}
		matchIdx[i] = struct{}{}

		for j := max(0, i-flags.Before); j < i; j++ {
			contextIdx[j] = struct{}{}
		}
		for j := i + 1; j <= min(len(lines)-1, i+flags.After); j++ {
			contextIdx[j] = struct{}{}
		}
	}
	return matchIdx, contextIdx
}

func readAllLines(r io.Reader) ([]string, error) {
	sc := bufio.NewScanner(r)
	const maxCap = 1024 * 1024
	sc.Buffer(make([]byte, 0, 64*1024), maxCap)

	var res []string
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res, sc.Err()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
