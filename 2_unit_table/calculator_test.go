package add_test

import (
	"testing"
	add "testing-demo/2_unit_table"
)

func TestAdd(t *testing.T) {
	// Test cases (Table)
	cases := []struct {
		name           string
		a, b, expected int
	}{
		{"Addition of positive numbers", 1, 2, 3},
		{"Addition of zeros", 0, 0, 0},
		{"Addition of positive and negative numbers", -1, 1, 0},
	}

	// Test each case
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := add.Add(c.a, c.b)
			if got != c.expected {
				t.Errorf("%s: Add(%d, %d) == %d, want %d", c.name, c.a, c.b, got, c.expected)
			}
		})
	}
}
