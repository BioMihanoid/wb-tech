package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
		f        flag
	}{
		{
			"default",
			[]string{"c", "a", "b"},
			[]string{"a", "b", "c"},
			flag{
				-1,
				false,
				false,
				false,
				"",
				"",
			},
		},
		{
			"sort reverse",
			[]string{"a", "b", "c"},
			[]string{"c", "b", "a"},
			flag{
				-1,
				false,
				true,
				false,
				"",
				"",
			},
		},
		{
			"sort unique",
			[]string{"a", "a", "c"},
			[]string{"a", "c"},
			flag{
				-1,
				false,
				false,
				true,
				"",
				"",
			},
		},
		{
			"sort numeric",
			[]string{"3", "a", "1", "2"},
			[]string{"1", "2", "3", "a"},
			flag{
				-1,
				true,
				false,
				false,
				"",
				"",
			},
		},
		{
			"sort by column",
			[]string{"x 9", "x 1", "x 5"},
			[]string{"x 1", "x 5", "x 9"},
			flag{
				2,
				false,
				false,
				false,
				"",
				"",
			},
		},
		{
			"multiply flags",
			[]string{"1", "3", "2", "4", "1"},
			[]string{"4", "3", "2", "1"},
			flag{
				-1,
				true,
				true,
				true,
				"",
				"",
			},
		},
	}

	t.Parallel()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := sortLines(tc.input, &tc.f)

			assert.Equal(t, tc.expected, res)
		})
	}
}
