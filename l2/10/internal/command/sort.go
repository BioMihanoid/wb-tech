package command

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"wb-tech/l2/10/pkg/files"

	"github.com/spf13/cobra"
)

type flag struct {
	Column         int
	Numeric        bool
	Reverse        bool
	Unique         bool
	InputPath      string
	OutputPath     string
	Month          bool
	IgnoreTrailing bool
	Check          bool
	HumanNumeric   bool
}

func sortCmd(cmd *cobra.Command, args []string) {
	f := &flag{}

	f.Column, _ = cmd.Flags().GetInt("column")
	f.Numeric, _ = cmd.Flags().GetBool("numeric")
	f.Reverse, _ = cmd.Flags().GetBool("reverse")
	f.Unique, _ = cmd.Flags().GetBool("unique")

	f.InputPath, _ = cmd.Flags().GetString("input")
	f.OutputPath, _ = cmd.Flags().GetString("output")

	f.Month, _ = cmd.Flags().GetBool("month")
	f.IgnoreTrailing, _ = cmd.Flags().GetBool("ignore-trailing")
	f.Check, _ = cmd.Flags().GetBool("check")
	f.HumanNumeric, _ = cmd.Flags().GetBool("human-numeric")

	var lines []string
	if f.InputPath != "" {
		lines = files.ReadLinesFromFile(f.InputPath)
	} else {
		lines = args
	}

	if f.Check {
		if !isSorted(lines) {
			fmt.Println("the lines are not sorted")
		} else {
			fmt.Println("the lines are sorted")
		}
		return
	}

	res := sortLines(lines, f)

	if f.OutputPath != "" {
		files.WriteLinesToFile(res, f.OutputPath)
	} else {
		for _, line := range res {
			fmt.Println(line)
		}
	}
}

func sortLines(lines []string, f *flag) []string {
	res := make([]string, len(lines))
	copy(res, lines)

	if f.Column == -1 && !f.Numeric && !f.Reverse && !f.Unique && !f.IgnoreTrailing && !f.Check && !f.Month &&
		!f.HumanNumeric {
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
	}

	if f.IgnoreTrailing {
		for i := range res {
			res[i] = strings.TrimRight(res[i], " ")
		}
	}

	if f.Column > 0 || f.Numeric {
		res = sortColumnNumeric(res, f.Column, f.Numeric)
	}

	if f.Reverse {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}

	if f.Unique {
		res = removeDuplicates(res)
	}

	if f.Month {
		res = sortMonth(res)
	}

	if f.HumanNumeric {
		sort.Slice(res, func(i, j int) bool {
			return parseHumanSize(res[i]) < parseHumanSize(res[j])
		})
	}

	return res
}

func sortColumnNumeric(lines []string, column int, numeric bool) []string {
	res := make([]string, len(lines))
	copy(res, lines)
	sort.Slice(res, func(i, j int) bool {
		var valI, valJ string

		if column > 0 {
			fieldsI := strings.Fields(res[i])
			fieldsJ := strings.Fields(res[j])
			if len(fieldsI) >= column {
				valI = fieldsI[column-1]
			}
			if len(fieldsJ) >= column {
				valJ = fieldsJ[column-1]
			}
		} else {
			valI = res[i]
			valJ = res[j]
		}

		if numeric {
			n1, err1 := strconv.Atoi(valI)
			n2, err2 := strconv.Atoi(valJ)
			if err1 == nil && err2 == nil {
				return n1 < n2
			}
			if err1 == nil {
				return true
			}
			if err2 == nil {
				return false
			}
		}

		return valI < valJ
	})

	return res
}

func removeDuplicates(lines []string) []string {
	m := make(map[string]struct{})
	res := make([]string, 0, len(lines))

	for _, v := range lines {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			res = append(res, v)
		}
	}

	return res
}

func isSorted(lines []string) bool {
	res := make([]string, len(lines))
	copy(res, lines)
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	for i := 0; i < len(res); i++ {
		if lines[i] != res[i] {
			return false
		}
	}
	return true
}

func sortMonth(lines []string) []string {
	res := make([]string, len(lines))
	copy(res, lines)

	sort.Slice(res, func(i, j int) bool {
		monthOrder := map[string]int{
			"jan": 1, "feb": 2, "mar": 3, "apr": 4,
			"may": 5, "jun": 6, "jul": 7, "aug": 8,
			"sep": 9, "oct": 10, "nov": 11, "dec": 12,
		}

		getMonthValue := func(s string) int {
			fields := strings.Fields(s)
			if len(fields) == 0 {
				return 13
			}
			month := strings.ToLower(fields[0])
			if val, ok := monthOrder[month]; ok {
				return val
			}
			return 13
		}

		return getMonthValue(res[i]) < getMonthValue(res[j])
	})

	return res
}

func parseHumanSize(s string) float64 {
	var suffixMultipliers = map[string]float64{
		"":  1,
		"K": 1e3,
		"M": 1e6,
		"G": 1e9,
		"T": 1e12,
		"P": 1e15,
		"E": 1e18,
	}

	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	s = strings.ToUpper(s)
	last := s[len(s)-1]
	numPart := s
	multiplier := 1.0

	if last >= 'A' && last <= 'Z' {
		numPart = s[:len(s)-1]
		if m, ok := suffixMultipliers[string(last)]; ok {
			multiplier = m
		}
	}

	num, err := strconv.ParseFloat(numPart, 64)
	if err != nil {
		return math.MaxFloat64
	}

	return num * multiplier
}
