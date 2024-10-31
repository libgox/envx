package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInt64(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected int64
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_INT64", "123", 123, true, false},
		{"EXISTING_NEGATIVE_INT64", "-456", -456, true, false},
		{"MAX_INT64", "9223372036854775807", 9223372036854775807, true, false},   // Maximum int64 value
		{"MIN_INT64", "-9223372036854775808", -9223372036854775808, true, false}, // Minimum int64 value
		{"OVERFLOW_INT64", "9223372036854775808", 0, true, true},                 // Overflow case for int64
		{"UNDERFLOW_INT64", "-9223372036854775809", 0, true, true},               // Underflow case for int64
		{"INVALID_INT64", "notanint", 0, true, true},                             // Non-integer string
		{"MISSING_INT64", "", 0, false, true},                                    // Missing variable case
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetInt64(tc.key)
			if tc.hasError {
				assert.Error(t, err, "Expected an error for key: %s", tc.key)
			} else {
				assert.NoError(t, err, "Did not expect an error for key: %s", tc.key)
				assert.Equal(t, tc.expected, result, "GetInt64(%s) = %v; expected %v", tc.key, result, tc.expected)
			}

			os.Unsetenv(tc.key)
		})
	}
}
