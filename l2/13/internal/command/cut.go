package command

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type flags struct {
	Fields    string
	Delimiter string
	Separated bool
}

func parseFields(fields string) (map[int]struct{}, error) {
	parts := strings.Split(fields, ",")
	res := make(map[int]struct{})
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if strings.Contains(part, "-") {
			nPart := strings.Split(part, "-")
			if len(nPart) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}
			start, err := strconv.Atoi(nPart[0])
			if err != nil || start <= 0 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}
			end, err := strconv.Atoi(nPart[1])
			if err != nil || end <= 0 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}
			if end < start {
				return nil, fmt.Errorf("invalid range (end < start): %s", part)
			}
			for i := start; i <= end; i++ {
				res[i] = struct{}{}
			}
		} else {
			num, err := strconv.Atoi(part)
			if err != nil || num <= 0 {
				return nil, fmt.Errorf("invalid fields: %s", fields)
			}
			res[num] = struct{}{}
		}
	}

	return res, nil
}

func cutCmd(flag *flags, r io.Reader, w io.Writer) {
	if flag.Delimiter == "" {
		flag.Delimiter = "\t"
	}

	neededColumns, err := parseFields(flag.Fields)
	if err != nil {
		_, _ = fmt.Fprintln(w, err)
		return
	}

	scanner := bufio.NewScanner(r)
	writer := bufio.NewWriter(w)

	for scanner.Scan() {
		line := scanner.Text()

		if flag.Separated && !strings.Contains(line, flag.Delimiter) {
			continue
		}

		parts := strings.Split(line, flag.Delimiter)
		out := make([]string, 0, len(parts))

		for idx, part := range parts {
			if _, ok := neededColumns[idx+1]; ok {
				out = append(out, part)
			}
		}

		_, _ = fmt.Fprintln(writer, strings.Join(out, flag.Delimiter))
		_ = writer.Flush()
	}

	if err = scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(w, "Read error:", err)
	}
}
