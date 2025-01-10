package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUint8(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected uint8
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_UINT8", "123", 123, true, false},
		{"MAX_UINT8", "255", 255, true, false},       // Maximum uint8 value
		{"MIN_UINT8", "0", 0, true, false},           // Minimum uint8 value
		{"OVERFLOW_UINT8", "256", 0, true, true},     // Overflow case for uint8
		{"INVALID_UINT8", "notanint", 0, true, true}, // Non-integer string
		{"MISSING_UINT8", "", 0, false, true},        // Missing variable case
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetUint8(tc.key)
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
