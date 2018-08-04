package contiguousArray

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
		{"2", 2, []int{0, 1}},
		{"4", 4, []int{0, 1, 1, 0, 1}},
		{"6", 6, []int{1, 1, 1, 0, 0, 1, 0, 1, 1}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, findMaxLength(c.input))
		})
	}
}
