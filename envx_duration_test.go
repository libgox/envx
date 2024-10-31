package envx

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDuration(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected time.Duration
		setEnv   bool
		hasError bool
	}{
		{"EXISTING_DURATION_MS", "500ms", 500 * time.Millisecond, true, false},
		{"EXISTING_DURATION_H", "1h", time.Hour, true, false},
		{"EXISTING_DURATION_COMPLEX", "2h30m", 2*time.Hour + 30*time.Minute, true, false},
		{"INVALID_DURATION", "notaduration", 0, true, true},
		{"MISSING_DURATION", "", 0, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			result, err := GetDuration(tc.key)
			if tc.hasError {
				assert.Error(t, err, "Expected an error for key: %s", tc.key)
			} else {
				assert.NoError(t, err, "Did not expect an error for key: %s", tc.key)
				assert.Equal(t, tc.expected, result, "GetDuration(%s) = %v; expected %v", tc.key, result, tc.expected)
			}

			os.Unsetenv(tc.key)
		})
	}
}
