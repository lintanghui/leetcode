// Source: https://leetcode.com/problems/longest-increasing-subsequence/
// Author: Lin Tanghui
// Date  : 2015/11/5

// Given an unsorted array of integers,
//  find the length of longest increasing subsequence.
// For example,
// Given [10, 9, 2, 5, 3, 7, 101, 18],
// The longest increasing subsequence is [2, 3, 7, 101],
// therefore the length is 4. Note that there may be more than one LIS combination, it is only necessary for you to return the length.

package LongestIncSub

func lengthofLIS(num []int) (lenth int) {
    var temp []int
    temp = make([]int, 0, len(num))
    for i := 0; i < len(num); i++ {
        if len(temp) == 0 || temp[len(temp)-1] <= num[i] {
            temp = append(temp, num[i])
        } else {
            var low, high = 0, len(temp)
            for {
                if low >= high {
                    break
                }
                j := (low + high) / 2
                if temp[j] > num[i] {
                    high = j
                } else {
                    low = j
                }
            }
            temp[j] = num[i]
        }
    }
    return len(temp)
}
