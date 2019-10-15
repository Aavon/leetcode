package d2

import (
	"fmt"
	"strings"
	"testing"
)


type ListNode struct {
    Val int
    Next *ListNode
}

func (l *ListNode) String() string {
	idx := l
	vals := []string{}
	for ;idx != nil; {
		vals = append(vals, fmt.Sprintf("%d", idx.Val))
		idx = idx.Next
	}
	return strings.Join(vals, " --> ")
}

func (l *ListNode) Init(vals []int) {
	if len(vals) == 0 {
		return
	}
	l.Val = vals[0]
	idx := l
	for i, val := range vals {
		if i == 0 {
			continue
		}
		idx.Next = &ListNode{
			Val: val,
		}
		idx = idx.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result ListNode
		idx0   *ListNode = &result
		idx1   *ListNode = l1
		idx2   *ListNode = l2
		carry bool
	)
	for ;(idx1 != nil || idx2 != nil); {
		next := &ListNode{}
		if idx1 != nil {
			next.Val += idx1.Val
		}
		if idx2 != nil {
			next.Val += idx2.Val
		}
		if carry {
			next.Val += 1
		}
		if next.Val >= 10 {
			carry = true
		} else {
			carry = false
		}
		next.Val = next.Val % 10
		idx0.Next = next
		idx0 = idx0.Next
		if idx1 != nil {
			idx1 = idx1.Next
		}
		if idx2 != nil {
			idx2 = idx2.Next
		}
	}
	if carry {
		idx0.Next = &ListNode{
			Val: 1,
		}
	}
	return result.Next
}


func Test_0(t *testing.T) {
	l1 := &ListNode{}
	l2 := &ListNode{}
	l1.Init([]int{2,4,3})
	l2.Init([]int{5,6,4})
	fmt.Printf("%v\n%v\n", l1, l2)
	result := addTwoNumbers(l1, l2)
	fmt.Printf("result: %v\n", result)
	t.Log("start")
}