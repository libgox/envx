package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStrOr(t *testing.T) {
	tests := []struct {
		key          string
		value        string
		defaultValue string
		expected     string
		setEnv       bool
	}{
		{"EXISTING_VAR", "Hello", "Default", "Hello", true},
		{"MISSING_VAR", "", "Default", "Default", false},
		{"EMPTY_VAR", "", "Fallback", "Fallback", true},
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result := GetStrOr(tc.key, tc.defaultValue)
			assert.Equal(t, tc.expected, result, "GetStrOr(%s, %s) = %s; expected %s", tc.key, tc.defaultValue, result, tc.expected)

			os.Unsetenv(tc.key)
		})
	}
}
