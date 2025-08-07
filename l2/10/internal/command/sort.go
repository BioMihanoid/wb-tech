package command

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"wb-tech/l2/10/pkg/files"

	"github.com/spf13/cobra"
)

type flag struct {
	Column     int
	Numeric    bool
	Reverse    bool
	Unique     bool
	InputPath  string
	OutputPath string
}

func sortCmd(cmd *cobra.Command, args []string) {
	f := &flag{}

	f.Column, _ = cmd.Flags().GetInt("column")
	f.Numeric, _ = cmd.Flags().GetBool("numeric")
	f.Reverse, _ = cmd.Flags().GetBool("reverse")
	f.Unique, _ = cmd.Flags().GetBool("unique")

	f.InputPath, _ = cmd.Flags().GetString("input")
	f.OutputPath, _ = cmd.Flags().GetString("output")

	var lines []string
	if f.InputPath != "" {
		lines = files.ReadLinesFromFile(f.InputPath)
	} else {
		lines = args
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

	if f.Column == -1 && !f.Numeric && !f.Reverse && !f.Unique {
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
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
