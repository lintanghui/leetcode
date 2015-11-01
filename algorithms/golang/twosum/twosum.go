// Source: https://leetcode.com/problems/two-sum/
// Author: Lin Tanghui
// Date  : 2015/10/30

// Given an array of integers, find two numbers such that they add up to a specific target number.
// The function twoSum should return indices of the two numbers such that they add up to the target,
// where index1 must be less than index2.
// Please note that your returned answers (both index1 and index2) are not zero-based.
// You may assume that each input would have exactly one solution.
// Input: numbers={2, 7, 11, 15}, target=9
// Output: index1=1, index2=2

package twosum

import (
    "sort"
)

func TwoSum(array []int, target int) (int, int) {
    var arrayMap map[int]int
    arrayMap = make(map[int]int, 0)
    for index, num := range array {
        arrayMap[num] = index
    }

    sort.Ints(array)
    low, high := 0, len(array)-1
    for {
        if low == high {
            break
        }
        sum := array[low] + array[high]
        switch {
        case sum == target:
            low, high = arrayMap[array[low]], arrayMap[array[high]]
            if low > high {
                return high + 1, low + 1
            } else {
                return low + 1, high + 1
            }
        case sum < target:
            low++
        case sum > target:
            high--
        }
    }
    // if not find return 0,0
    return 0, 0
}
