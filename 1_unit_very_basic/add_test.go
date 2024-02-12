package add_test

import (
	"testing"
	add "testing-demo/1_unit_very_basic"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// define inputs
	inputA := 1
	inputB := 2

	// define expected result
	expected := 3

	// perform test
	actual := add.Add(inputA, inputB)

	// assert that the actual result is equal to expected
	assert.Equal(t, expected, actual)
}
