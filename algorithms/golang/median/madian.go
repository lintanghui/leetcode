// Source: https://leetcode.com/problems/median-of-two-sorted-arrays/
// Author: Lin Tanghui
// Date  : 2015/10/31

// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).

package median

import (
	"fmt"
)

func FindSortedArrays(a1, a2 []int, size1, size2 int) float32 {
	var (
		array []int = make([]int, 0)
		i, j  int
	)
	if size1 == 0 {
		return median(a2, size2)
	}
	if size2 == 0 {
		return median(a1, size1)
	}
	for {
		if a1[i] < a2[j] {

			array = append(array, a1[i])
			i++
			if i >= size1 {

				for j < size2 {
					array = append(array, a2[j])
					j++
				}
				break
			}
		} else {
			array = append(array, a2[j])
			j++
			if j >= size2 {
				for i < size1 {
					array = append(array, a1[i])
					i++
				}
				break
			}

		}

	}

	return median(array, len(array))
}

func median(array []int, lenth int) (num float32) {
	mod := lenth % 2
	mid := lenth / 2
	fmt.Println(array)
	if mod == 0 {
		num = float32(array[mid]+array[mid-1]) / 2.0

	} else {
		num = float32(array[mid])

	}
	return
}
