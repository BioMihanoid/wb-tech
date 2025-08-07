package pack

import (
	"errors"
	"strconv"
	"strings"
)

func UnpackingStr(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var result strings.Builder
	var prev rune
	escaped := false
	expectChar := true

	for _, r := range s {
		switch {
		case escaped:
			result.WriteRune(r)
			prev = r
			escaped = false
			expectChar = false
		case r == '\\':
			escaped = true
		case isDigit(r):
			if expectChar {
				return "", errors.New("string cannot start with digit or contain digit without preceding character")
			}
			repeat, _ := strconv.Atoi(string(r))
			result.WriteString(strings.Repeat(string(prev), repeat-1))
			expectChar = true
		default:
			result.WriteRune(r)
			prev = r
			expectChar = false
		}
	}

	if escaped {
		return "", errors.New("string ends with single backslash")
	}

	return result.String(), nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
