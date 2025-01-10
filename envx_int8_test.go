package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInt8(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected int8
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_INT8", "123", 123, true, false},
		{"EXISTING_NEGATIVE_INT8", "-123", -123, true, false},
		{"MAX_INT8", "127", 127, true, false},       // Maximum int8 value
		{"MIN_INT8", "-128", -128, true, false},     // Minimum int8 value
		{"OVERFLOW_INT8", "128", 0, true, true},     // Overflow case for int8
		{"UNDERFLOW_INT8", "-129", 0, true, true},   // Underflow case for int8
		{"INVALID_INT8", "notanint", 0, true, true}, // Non-integer string
		{"MISSING_INT8", "", 0, false, true},        // Missing variable case
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetInt8(tc.key)
			if tc.hasError {
				assert.Error(t, err, "Expected an error for key: %s", tc.key)
			} else {
				assert.NoError(t, err, "Did not expect an error for key: %s", tc.key)
				assert.Equal(t, tc.expected, result, "GetInt8(%s) = %v; expected %v", tc.key, result, tc.expected)
			}

			os.Unsetenv(tc.key)
		})
	}
}
