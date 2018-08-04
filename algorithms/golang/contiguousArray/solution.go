package contiguousArray

func findMaxLength(nums []int) int {
	for i, v := range nums {
		if v == 0 {
			nums[i] = -1
		}
	}
	sums := make(map[int]int, 0)
	sums[0] = -1
	var sum int
	var max int
	for i, v := range nums {
		sum += v
		if idx, ok := sums[sum]; ok {
			max = findMax(max, i-idx)
		} else {
			sums[sum] = i
		}
	}
	return max
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
