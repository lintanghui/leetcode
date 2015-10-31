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
		substring [256]rune
		index     = 0
	)
	for _, ch := range s {
		if substring[index] != ch {
			substring[index] = ch
			index++
		} else {
			index = 0
			substring[index] = ch
		}
	}
	return len(substring)
}
