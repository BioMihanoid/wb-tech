package command

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var sample = []string{
	"alpha",
	"Beta",
	"gamma",
	"betamax",
	"delta",
}

func runGrep(t *testing.T, flags *grepFlags, pattern string, lines []string) string {
	t.Helper()
	m, err := buildMatcher(pattern, flags.Fixed, flags.IgnoreCase)
	require.NoError(t, err)

	var out bytes.Buffer
	printGrep(flags, lines, m, "", &out)
	return out.String()
}

func TestGrep_SimpleMatch(t *testing.T) {
	flags := &grepFlags{}
	got := runGrep(t, flags, "beta", sample)
	require.Equal(t, "betamax\n", got)
}

func TestGrep_IgnoreCase(t *testing.T) {
	flags := &grepFlags{IgnoreCase: true}
	got := runGrep(t, flags, "beta", sample)
	lines := strings.Split(strings.TrimSpace(got), "\n")
	require.ElementsMatch(t, []string{"Beta", "betamax"}, lines)
}

func TestGrep_InvertMatch(t *testing.T) {
	flags := &grepFlags{InvertMatch: true}
	got := runGrep(t, flags, "beta", sample)
	lines := strings.Split(strings.TrimSpace(got), "\n")
	require.ElementsMatch(t, []string{"alpha", "Beta", "gamma", "delta"}, lines)
}

func TestGrep_FixedStrings(t *testing.T) {
	flags := &grepFlags{Fixed: true}
	got := runGrep(t, flags, "Beta", sample)
	require.Equal(t, "Beta\n", got)

	flags = &grepFlags{Fixed: true, IgnoreCase: true}
	got = runGrep(t, flags, "beta", sample)
	lines := strings.Split(strings.TrimSpace(got), "\n")
	require.ElementsMatch(t, []string{"Beta", "betamax"}, lines)
}

func TestGrep_CountOnly(t *testing.T) {
	flags := &grepFlags{Count: true, IgnoreCase: true}
	got := runGrep(t, flags, "beta", sample)
	require.Equal(t, "2\n", got) // "Beta" и "betamax"
}

func TestGrep_ContextAfterBefore(t *testing.T) {
	flags := &grepFlags{IgnoreCase: true, Before: 1, After: 1}
	got := runGrep(t, flags, "gamma", sample)
	expect := "Beta\ngamma\nbetamax\n"
	require.Equal(t, expect, got)
}

func TestGrep_ContextC(t *testing.T) {
	flags := &grepFlags{IgnoreCase: true, Context: 1}
	normalizeContextFlags(flags)
	got := runGrep(t, flags, "beta", sample)

	expect := "alpha\nBeta\ngamma\nbetamax\ndelta\n"
	require.Equal(t, expect, got)
}

func TestGrep_LineNumbers(t *testing.T) {
	flags := &grepFlags{LineNumber: true, IgnoreCase: true, Fixed: true}
	got := runGrep(t, flags, "alpha", sample)
	require.Equal(t, "1:alpha\n", got)
}

func TestGrep_InvalidRegex(t *testing.T) {
	_, err := buildMatcher("([a-", false, false)
	require.Error(t, err)
}
