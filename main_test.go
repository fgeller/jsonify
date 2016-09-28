package main

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		schema  string
		json    string
		isValid bool
	}{
		{
			schema:  `{"type": "object"}`,
			json:    `[1, "2", 3]`,
			isValid: false,
		},
		{
			schema:  `{"type": "object"}`,
			json:    `{"name": "john"}`,
			isValid: true,
		},
	}

	for i, test := range tests {
		actual := validate(test.schema, test.json)
		actualIsValid := actual == nil
		if (test.isValid && !actualIsValid) || (!test.isValid && actualIsValid) {
			t.Errorf("Test %d: expected [%v], but received %T\n", i, test.isValid, actual)
		}
	}
}
