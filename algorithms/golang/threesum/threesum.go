// Source: https://leetcode.com/problems/3sum/
// Author: Lin Tanghui
// Date  : 2015/10/31

// Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.

package threesum

import (
    "fmt"
    "sort"
)

type List []int

func threesum(list []int) []List {
    var ansList []List
    lenth := len(list)
    sort.Ints(list)
    for i := 0; i < lenth-2; i++ {

        low, high := i+1, lenth-1
        sum := 0 - list[i]
        for {
            if low >= high {
                break
            }

            switch {
            case list[low]+list[high] == sum:
                var ans List = []int{list[i], list[low], list[high]}
                fmt.Println(ans)
                ansList = append(ansList, ans)
                low++
                high = high - 1
            case list[low]+list[high] < sum:
                low++
            case list[low]+list[high] > sum:
                high = high - 1
            }

        }
    }
    return ansList
}
