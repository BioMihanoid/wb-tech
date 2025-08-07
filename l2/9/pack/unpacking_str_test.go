package pack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpackingStr(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
			wantErr:  false,
		},
		{
			name:     "no digits",
			input:    "abcd",
			expected: "abcd",
			wantErr:  false,
		},
		{
			name:     "simple repeat",
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
			wantErr:  false,
		},
		{
			name:     "only digits - invalid",
			input:    "45",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "escaped digits",
			input:    "qwe\\4\\5",
			expected: "qwe45",
			wantErr:  false,
		},
		{
			name:     "partial escaped",
			input:    "qwe\\45",
			expected: "qwe44444",
			wantErr:  false,
		},
		{
			name:     "ends with escape",
			input:    "abc\\",
			expected: "",
			wantErr:  true,
		},
	}

	t.Parallel()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := UnpackingStr(tc.input)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Empty(t, actual)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}
