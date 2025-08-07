package command

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type flag struct {
	Column  int
	Numeric bool
	Reverse bool
	Unique  bool
}

func sortCmd(cmd *cobra.Command, args []string) {
	fmt.Println("start sort")
	f := &flag{}

	f.Column, _ = cmd.Flags().GetInt("column")
	f.Numeric, _ = cmd.Flags().GetBool("numeric")
	f.Reverse, _ = cmd.Flags().GetBool("reverse")
	f.Unique, _ = cmd.Flags().GetBool("unique")

	res := sortLines(args, f)

	for _, v := range res {
		fmt.Println(v)
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
		sort.Slice(res, func(i, j int) bool {
			var valI, valJ string

			if f.Column > 0 {
				fieldsI := strings.Fields(res[i])
				fieldsJ := strings.Fields(res[j])
				if len(fieldsI) >= f.Column {
					valI = fieldsI[f.Column-1]
				}
				if len(fieldsJ) >= f.Column {
					valJ = fieldsJ[f.Column-1]
				}
			} else {
				valI = res[i]
				valJ = res[j]
			}

			if f.Numeric {
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

func removeDuplicates(input []string) []string {
	m := make(map[string]struct{})
	res := make([]string, 0, len(input))

	for _, v := range input {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			res = append(res, v)
		}
	}

	return res
}
