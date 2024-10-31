package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBool(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected bool
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_TRUE", "true", true, true, false},
		{"EXISTING_FALSE", "false", false, true, false},
		{"INVALID_BOOL", "notabool", false, true, true},
		{"MISSING_VAR", "", false, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetBool(tc.key)
			if tc.hasError {
				assert.Error(t, err, "Expected an error for key: %s", tc.key)
			} else {
				assert.NoError(t, err, "Did not expect an error for key: %s", tc.key)
				assert.Equal(t, tc.expected, result, "GetBool(%s) = %v; expected %v", tc.key, result, tc.expected)
			}

			os.Unsetenv(tc.key)
		})
	}
}
