package divideTwoIntegers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name   string
		expect int
		input  []int
	}{
		{"10-3", 3, []int{10, 3}},
		{"1-1", 1, []int{1, 1}},
		{"10- -3", -3, []int{10, -3}},
		{"-10- -3", 3, []int{-10, -3}},
		{"10-4", 2, []int{10, 4}},
		{"9-3", 3, []int{9, 3}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, divide(c.input[0], c.input[1]))
		})
	}
}
