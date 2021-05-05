package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func lton(l *ListNode) int {
	n := 0
	times := 1
	for l != nil {
		n = n + l.Val*times
		l = l.Next
		times *= 10
	}
	return n
}

func ntol(n int) *ListNode {
	l := &ListNode{Val: n % 10}
	t := l
	for n >= 10 {
		n = n / 10
		t.Next = &ListNode{Val: n % 10}
		t = t.Next
	}
	return l
}

func main() {
	l1 := ntol(243)
	l2 := ntol(564)
	fmt.Println(lton(addTwoNumber(l1, l2)))
}
