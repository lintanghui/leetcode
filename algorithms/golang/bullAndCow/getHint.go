// Source: https://leetcode.com/problems/bulls-and-cows/
// Author: Lin Tanghui
// Date  : 2015/11/4

// You are playing the following Bulls and Cows game with your friend: You write a 4-digit secret number and ask your friend to guess it.
// Each time your friend guesses a number, you give a hint. The hint tells your friend how many digits are in the correct positions (called "bulls")
// and how many digits are in the wrong positions (called "cows").Your friend will use those hints to find out the secret number.
package bullAndCow

import (
    "fmt"
)

func getHint(secret, guest string) (ans string) {
    var (
        vsec      [10]int
        vgue      [10]int
        bull, cow = 0, 0
    )

    if len(secret) != len(guest) {
        return "0A0B"
    }
    for i := 0; i < len(secret); i++ {
        if secret[i] == guest[i] {
            bull++
        } else {
            vgue[guest[i]-'0']++
            vsec[secret[i]-'0']++
        }

    }
    for i := 0; i < len(vsec); i++ {
        cow += min(vgue[i], vsec[i])
    }
    return fmt.Sprintf("%dA%dB", bull, cow)
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}
