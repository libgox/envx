package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInt16(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected int16
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_INT16", "123", 123, true, false},
		{"EXISTING_NEGATIVE_INT16", "-456", -456, true, false},
		{"MAX_INT16", "32767", 32767, true, false},   // Maximum int16 value
		{"MIN_INT16", "-32768", -32768, true, false}, // Minimum int16 value
		{"OVERFLOW_INT16", "32768", 0, true, true},   // Overflow case for int16
		{"UNDERFLOW_INT16", "-32769", 0, true, true}, // Underflow case for int16
		{"INVALID_INT16", "notanint", 0, true, true}, // Non-integer string
		{"MISSING_INT16", "", 0, false, true},        // Missing variable case
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetInt16(tc.key)
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
