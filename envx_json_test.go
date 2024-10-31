package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJson(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct {
		key      string
		value    string
		target   any
		expected any
		setEnv   bool
		hasError bool
	}{
		{
			key:      "PERSON_JSON",
			value:    `{"name": "John", "age": 30}`,
			target:   &Person{},
			expected: &Person{Name: "John", Age: 30},
			setEnv:   true,
			hasError: false,
		},
		{
			key:      "MAP_JSON",
			value:    `{"key1": "value1", "key2": "value2"}`,
			target:   &map[string]string{},
			expected: &map[string]string{"key1": "value1", "key2": "value2"},
			setEnv:   true,
			hasError: false,
		},
		{
			key:      "INVALID_JSON",
			value:    `{"name": "Alice", "age": }`,
			target:   &Person{},
			expected: &Person{},
			setEnv:   true,
			hasError: true,
		},
		{
			key:      "MISSING_JSON",
			value:    "",
			target:   &Person{},
			expected: &Person{},
			setEnv:   false,
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv(tc.key, tc.value)
			}

			err := GetJson(tc.key, tc.target)
			if tc.hasError {
				assert.Error(t, err, "Expected an error for key: %s", tc.key)
			} else {
				assert.NoError(t, err, "Did not expect an error for key: %s", tc.key)
				assert.Equal(t, tc.expected, tc.target, "GetJSON(%s) = %v; expected %v", tc.key, tc.target, tc.expected)
			}

			os.Unsetenv(tc.key)
		})
	}
}
