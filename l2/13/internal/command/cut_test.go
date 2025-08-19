package command

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCutCmd_Basic(t *testing.T) {
	input := `one two three four
a b c d
`
	expected := `one three
a c
`

	flag := &flags{
		Fields:    "1,3",
		Delimiter: " ",
	}

	var out bytes.Buffer
	cutCmd(flag, strings.NewReader(input), &out)

	assert.Equal(t, expected, out.String())
}

func TestCutCmd_Range(t *testing.T) {
	input := `one two three four
a b c d
`
	expected := `one two three
a b c
`

	flag := &flags{
		Fields:    "1-3",
		Delimiter: " ",
	}

	var out bytes.Buffer
	cutCmd(flag, strings.NewReader(input), &out)

	assert.Equal(t, expected, out.String())
}

func TestCutCmd_SeparatedFlag(t *testing.T) {
	input := `one two three
notabhere
a b c
`
	expected := `one two
a b
`

	flag := &flags{
		Fields:    "1,2",
		Delimiter: " ",
		Separated: true,
	}

	var out bytes.Buffer
	cutCmd(flag, strings.NewReader(input), &out)

	assert.Equal(t, expected, out.String())
}

func TestCutCmd_InvalidFields(t *testing.T) {
	input := `one two three`
	flag := &flags{
		Fields:    "abc",
		Delimiter: " ",
	}

	var out bytes.Buffer
	cutCmd(flag, strings.NewReader(input), &out)

	assert.Contains(t, out.String(), "invalid fields")
}

func TestCutCmd_EmptyDelimiterDefaultsToTab(t *testing.T) {
	input := "a\tb\tc\nx\ty\tz\n"
	expected := "a\tc\nx\tz\n"

	flag := &flags{
		Fields:    "1,3",
		Delimiter: "",
	}

	var out bytes.Buffer
	cutCmd(flag, strings.NewReader(input), &out)

	assert.Equal(t, expected, out.String())
}
