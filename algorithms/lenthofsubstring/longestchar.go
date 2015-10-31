// Source: https://leetcode.com/problems/two-sum/
// Author: Lin Tanghui
// Date  : 2015/10/30

// Given a string, find the length of the longest substring without repeating characters.
// For example, the longest substring without repeating letters
// for "abcabcbb" is "abc", which the length is 3.
// For "bbbbb" the longest substring is "b", with the length of 1.

package lenthofsubstring

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	var (
		substring       map[rune]int = make(map[rune]int, 0)
		m, longest, num int
	)
	for index, ch := range s {
		num, _ = substring[ch]
		m = max(num, m)
		substring[ch] = index
		longest = max(longest, index-m)
	}
	return longest

}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}

}
