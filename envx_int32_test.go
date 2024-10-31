package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInt32(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected int32
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_INT32", "123", 123, true, false},
		{"EXISTING_NEGATIVE_INT32", "-456", -456, true, false},
		{"MAX_INT32", "2147483647", 2147483647, true, false},   // Maximum int32 value
		{"MIN_INT32", "-2147483648", -2147483648, true, false}, // Minimum int32 value
		{"OVERFLOW_INT32", "2147483648", 0, true, true},        // Overflow case for int32
		{"UNDERFLOW_INT32", "-2147483649", 0, true, true},      // Underflow case for int32
		{"INVALID_INT32", "notanint", 0, true, true},           // Non-integer string
		{"MISSING_INT32", "", 0, false, true},                  // Missing variable case
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetInt32(tc.key)
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
