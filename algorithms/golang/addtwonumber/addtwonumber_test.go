package addtwonumber

import (
    "testing"
)

func TestAddTwoNumber(t *testing.T) {
    var l1 = structList{1, &structList{9, &structList{3, nil}}}
    var l2 = structList{2, &structList{3, &structList{4, nil}}}
    AddTwoNumbers(&l1, &l2)
    var l3 = structList{1, &structList{8, &structList{9, &structList{9, &structList{9, nil}}}}}
    var l4 = structList{6, &structList{5, &structList{4, nil}}}
    AddTwoNumbers(&l3, &l4)
}
