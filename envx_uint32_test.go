package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUint32(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected uint32
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_UINT32", "123", 123, true, false},
		{"MAX_UINT32", "4294967295", 4294967295, true, false}, // Maximum uint32 value
		{"MIN_UINT32", "0", 0, true, false},                   // Minimum uint32 value
		{"OVERFLOW_UINT32", "4294967296", 0, true, true},      // Overflow case for uint32
		{"INVALID_UINT32", "notanint", 0, true, true},         // Non-integer string
		{"MISSING_UINT32", "", 0, false, true},                // Missing variable case
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetUint32(tc.key)
			if tc.hasError {
				assert.Error(t, err, "Expected an error for key: %s", tc.key)
			} else {
				assert.NoError(t, err, "Did not expect an error for key: %s", tc.key)
				assert.Equal(t, tc.expected, result, "GetInt32(%s) = %v; expected %v", tc.key, result, tc.expected)
			}

			os.Unsetenv(tc.key)
		})
	}
}
