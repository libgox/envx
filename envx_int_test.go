package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInt(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected int
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_INT", "123", 123, true, false},
		{"EXISTING_NEGATIVE_INT", "-456", -456, true, false},
		{"INVALID_INT", "notanint", 0, true, true},
		{"MISSING_INT", "", 0, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetInt(tc.key)
			if tc.hasError {
				assert.Error(t, err, "Expected an error for key: %s", tc.key)
			} else {
				assert.NoError(t, err, "Did not expect an error for key: %s", tc.key)
				assert.Equal(t, tc.expected, result, "GetInt(%s) = %v; expected %v", tc.key, result, tc.expected)
			}

			os.Unsetenv(tc.key)
		})
	}
}
