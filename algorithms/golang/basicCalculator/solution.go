package basicCalculator

import (
	"fmt"
)

func calculate(s string) int {
	var temp int
	var nums []int
	var ops []rune
	for i, c := range s {
		if c == ' ' {
			continue
		}
		if c >= '0' && c <= '9' {
			temp = temp*10 + (int(c) - 48)
			if i == len(s)-1 {
				nums = append(nums, temp)
			}
		} else {
			nums = append(nums, temp)
			temp = 0
		}

		if c == '+' || c == '-' || c == '*' || c == '/' {
			ops = append(ops, c)
		}
	}
	if temp != 0 {
		nums = append(nums, temp)
	}
	fmt.Println(nums, ops)
	var results []int
	var rsop []rune
	results = append(results, nums[0])
	for i, op := range ops {
		if op == '*' || op == '/' {
			pre := results[len(results)-1]
			if op == '*' {
				results[len(results)-1] = pre * nums[i+1]
			} else {
				results[len(results)-1] = pre / nums[i+1]
			}
			fmt.Println("tmp", results)
			continue
		} else {
			results = append(results, nums[i+1])
		}
		rsop = append(rsop, op)
	}
	fmt.Println("rs", results, rsop)
	var result = results[0]
	for i, op := range rsop {
		if op == '+' {
			result = result + results[i+1]
		} else {
			result = result - results[i+1]
		}
	}
	return result
}
