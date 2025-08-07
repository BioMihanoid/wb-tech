package command

import (
	"fmt"
	"math"
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
				false,
				false,
				false,
				false,
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
				false,
				false,
				false,
				false,
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
				false,
				false,
				false,
				false,
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
				false,
				false,
				false,
				false,
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
				false,
				false,
				false,
				false,
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
				false,
				false,
				false,
				false,
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

func TestAddSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
		f        flag
	}{
		{
			"sort month",
			[]string{"feb", "nov", "aboba", "dec", "bob", "jan"},
			[]string{"jan", "feb", "nov", "dec", "aboba", "bob"},
			flag{
				-1,
				false,
				false,
				false,
				"",
				"",
				true,
				false,
				false,
				false,
			},
		},
		{
			"ignore trailing",
			[]string{"feb    ", "nov", "aboba        ", "dec", "bob      ", "jan           "},
			[]string{"jan", "feb", "nov", "dec", "aboba", "bob"},
			flag{
				-1,
				false,
				false,
				false,
				"",
				"",
				true,
				true,
				false,
				false,
			},
		},
		{
			"human numeric",
			[]string{"apple", "1K", "banana", "500"},
			[]string{"500", "1K", "apple", "banana"},
			flag{
				-1,
				false,
				false,
				false,
				"",
				"",
				false,
				false,
				false,
				true,
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

func TestIsSorted(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected bool
	}{
		{"already sorted", []string{"a", "b", "c"}, true},
		{"not sorted", []string{"c", "a", "b"}, false},
		{"empty", []string{}, true},
		{"single element", []string{"a"}, true},
		{"with duplicates", []string{"a", "a", "b"}, true},
	}

	t.Parallel()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isSorted(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestSortMonth(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			"mixed case months",
			[]string{"FEB", "jan", "Mar"},
			[]string{"jan", "FEB", "Mar"},
		},
		{
			"invalid months",
			[]string{"abc", "jan", "xyz"},
			[]string{"jan", "abc", "xyz"},
		},
	}

	t.Parallel()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := sortMonth(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestParseHumanSize(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected float64
	}{
		{"empty string", "", 0},
		{"invalid number", "abcK", math.MaxFloat64},
		{"negative number", "-1.5K", -1.5 * 1e3},
		{"unicode spaces", "　1K　", 1000},
	}

	t.Parallel()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := parseHumanSize(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestSortColumnNumeric_EdgeCases(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		column   int
		numeric  bool
		expected []string
	}{
		{
			"column out of bounds",
			[]string{"a b c", "x y z"},
			100,
			false,
			[]string{"a b c", "x y z"},
		},
		{
			"numeric with invalid numbers",
			[]string{"a 10", "b 2x", "c 5"},
			2,
			true,
			[]string{"c 5", "a 10", "b 2x"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := sortColumnNumeric(tc.input, tc.column, tc.numeric)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestRemoveDuplicates_EdgeCases(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"a", "a", "a"}, []string{"a"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			res := removeDuplicates(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
