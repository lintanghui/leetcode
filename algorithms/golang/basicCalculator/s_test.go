package basicCalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name   string
		expect int
		input  string
	}{
		{"1", 3, "1+1+2/2"},
		{"10", 10, "1+5+2*2"},
		{"10 ", 10, "1+ 5+2*2"},
		{"9", 9, "1*5+2*2"},
		{"1/1", 4, "1*5/2*2"},
		{"3/2", 1, " 3/2  "},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, calculate(c.input))
		})
	}
}
