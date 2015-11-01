// Source: https://leetcode.com/problems/add-two-numbers/
// Author: Lin Tanghui
// Date  : 2015/10/31

// You are given two linked lists representing two non-negative numbers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8

package addtwonumber

import (
    "fmt"
)

type structList struct {
    val  int
    next *structList
}

func AddTwoNumbers(l1, l2 *structList) *structList {
    switch {
    case l1 == nil && l2 == nil:
        return l1
    case l1 == nil:
        return l2
    case l2 == nil:
        return l1
    }
    var (
        head  = l1
        pre   = l1
        carry int
    )
    pre.val = add(l1.val, l2.val, &carry)
    head = pre

    l1 = l1.next
    l2 = l2.next
    for {
        if l1 == nil || l2 == nil {
            break
        }
        pre.next.val = add(l1.val, l2.val, &carry)
        pre = pre.next
        l1 = l1.next
        l2 = l2.next
    }
    if l2 != nil {
        pre.next = l2
    }

    for carry != 0 {
        if pre.next != nil {
            pre.next.val = add(pre.next.val, 0, &carry)
            pre = pre.next
        } else {
            pre.next = &structList{carry, nil}
            break
        }
    }
    printList(head)
    return head
}

func add(a, b int, carry *int) int {
    var sum = a + b + *carry
    *carry = sum / 10
    return sum % 10
}

func printList(l *structList) {
    for {
        if l != nil {
            fmt.Printf("%d-->", l.val)
            l = l.next
        } else {
            fmt.Println("")
            return
        }

    }

}
