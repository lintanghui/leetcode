package LongestIncSub

import (
    "testing"
)

func TestLIS(t *testing.T) {
    a := []int{1, 2, 3, 5, 9, 4}
    lenth := lengthofLIS(num)
    if lenth != 5 {
        t.Error(lenth)
    }
}
