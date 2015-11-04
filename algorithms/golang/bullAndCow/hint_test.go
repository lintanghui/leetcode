package bullAndCow

import (
    "testing"
)

func TestGetHint(t *testing.T) {
    ans := getHint("3456", "5673")
    if ans != "0A3B" {
        t.Error(ans)
    }
    ans = getHint("0123", "0132")
    if ans != "2A2B" {
        t.Error(ans)
    }
}
