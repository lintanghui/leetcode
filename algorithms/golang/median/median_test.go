package median

import (
	"testing"
)

func TestMedian(t *testing.T) {
	var (
		a1 = []int{1, 2, 3, 4, 7}
		a2 = []int{2, 3, 6, 8, 9}
		a3 = []int{1, 2, 3, 4, 7, 111}
		a4 = []int{2, 3, 6, 8, 9}
	)
	median := FindSortedArrays(a1, a2, 5, 5)
	if median != 3.5 {
		t.Error(median)
	}
	median = FindSortedArrays(a3, a4, 6, 5)
	if median != 4 {
		t.Error(median)
	}
}
