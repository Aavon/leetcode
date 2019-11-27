package d24

import (
	"fmt"
	"strings"
	"testing"
)

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (62.52%)
 * Likes:    328
 * Dislikes: 0
 * Total Accepted:    50.6K
 * Total Submissions: 80.6K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
 *
 *
 */

// @lc code=start
//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
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

func (l *ListNode) String() string {
	idx := l
	vals := []string{}
	for ;idx != nil; {
		vals = append(vals, fmt.Sprintf("%d", idx.Val))
		idx = idx.Next
	}
	return strings.Join(vals, " --> ")
}

func swapPairs(head *ListNode) *ListNode {
	var before *ListNode = nil
	var result *ListNode = head
	t := head
	for ;t != nil && t.Next != nil; {
		right := t.Next
		t.Next, t.Next.Next = t.Next.Next, t
		if before != nil {
			before.Next = right
		} else {
			result = right
		}
		before = right.Next
		t = before.Next
	}
	return result
}

func Test_a(t *testing.T) {
	l1 := &ListNode{}
	l1.Init([]int{5,6,4})
	fmt.Printf("%v\n", l1)
	l2 := swapPairs(l1)
	fmt.Printf("%v\n", l2)
}
