package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUint16(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected uint16
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_UINT16", "123", 123, true, false},
		{"MAX_UINT16", "65535", 65535, true, false},   // Maximum uint16 value
		{"MIN_UINT16", "0", 0, true, false},           // Minimum uint16 value
		{"OVERFLOW_UINT16", "65536", 0, true, true},   // Overflow case for uint16
		{"INVALID_UINT16", "notanint", 0, true, true}, // Non-integer string
		{"MISSING_UINT16", "", 0, false, true},        // Missing variable case
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetUint16(tc.key)
			if tc.hasError {
				assert.Error(t, err, "Expected an error for key: %s", tc.key)
			} else {
				assert.NoError(t, err, "Did not expect an error for key: %s", tc.key)
				assert.Equal(t, tc.expected, result, "GetInt16(%s) = %v; expected %v", tc.key, result, tc.expected)
			}

			os.Unsetenv(tc.key)
		})
	}
}
