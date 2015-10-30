package twosum

import (
    "testing"
)

func TestTwoSum(t *testing.T) {
    var (
        array1  = []int{1, 2, 5, 7, 8, 3}
        array2  = []int{13, 6, 8, 2, 4, 5}
        target1 = 11
        target2 = 14
    )
    low, high := TwoSum(array1, target1)
    if low != 5 || high != 6 {
        t.Errorf("two sum error %d,%d", low, high)
    }

    low, high = TwoSum(array2, target2)
    if low != 2 || high != 3 {
        t.Errorf("two sum error %d,%d", low, high)
    }
}
