package threesum

import (
    "testing"
)

func TestThreeSum(t *testing.T) {
    data := []int{0, 1, -2, -3, 2, 3, -1}
    ans := threesum(data)
    t.Log(ans)

}
