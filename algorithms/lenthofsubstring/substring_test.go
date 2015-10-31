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
	if 6 != lenth {
		t.Error(lenth)
	}
	lenth = lengthOfLongestSubstring(str2)
	if 4 != lenth {
		t.Error(lenth)
	}
	lenth = lengthOfLongestSubstring("")
	if 0 != lenth {
		t.Error(lenth)
	}
}
