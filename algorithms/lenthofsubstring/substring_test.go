package lenthofsubstring

import (
	"testing"
)

func TestLenth(t *testing.T) {
	var (
		str1 = "aaabvcsdaa"
		str2 = "abccabcd"
	)
	lenth := lengthOfLongestSubstring(str1)
	if 7 != lenth {
		t.Log(lenth)
	}
	lenth = lengthOfLongestSubstring(str2)
	if 5 != lenth {
		t.Log(lenth)
	}
	lenth = lengthOfLongestSubstring("")
	if 0 != lenth {
		t.Log(lenth)
	}
}
