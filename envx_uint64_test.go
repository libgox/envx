package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUint64(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected uint64
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_UINT64", "123", 123, true, false},
		{"MAX_UINT64", "18446744073709551615", 18446744073709551615, true, false}, // Maximum uint64 value
		{"MIN_UINT64", "0", 0, true, false},                                       // Minimum uint64 value
		{"OVERFLOW_UINT64", "18446744073709551616", 0, true, true},                // Overflow case for uint64
		{"INVALID_UINT64", "notanint", 0, true, true},                             // Non-integer string
		{"MISSING_UINT64", "", 0, false, true},                                    // Missing variable case
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetUint64(tc.key)
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
