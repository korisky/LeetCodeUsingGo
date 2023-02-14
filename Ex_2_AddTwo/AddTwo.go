package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l1_3 := ListNode{Val: 9}
	l1_2 := ListNode{Val: 9, Next: &l1_3}
	l1_1 := ListNode{Val: 9, Next: &l1_2}

	l2_3 := ListNode{Val: 9}
	l2_2 := ListNode{Val: 9, Next: &l2_3}
	l2_1 := ListNode{Val: 9, Next: &l2_2}

	head := addTwoNumbers_Advanced(&l1_1, &l2_1)
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	sumHead := new(ListNode)
	curHead := sumHead

	for l1 != nil && l2 != nil {
		// calculating
		curSum := l1.Val + l2.Val + carry
		nextNode := ListNode{
			Val:  curSum % 10,
			Next: nil,
		}
		// moving next
		curHead.Next = &nextNode
		curHead = curHead.Next
		l1 = l1.Next
		l2 = l2.Next
		carry = curSum / 10
	}

	for l1 != nil {
		curSum := l1.Val + carry
		nextNode := ListNode{
			Val:  curSum % 10,
			Next: nil,
		}
		// moving next
		curHead.Next = &nextNode
		curHead = curHead.Next
		l1 = l1.Next
		carry = curSum / 10
	}

	for l2 != nil {
		curSum := l2.Val + carry
		nextNode := ListNode{
			Val:  curSum % 10,
			Next: nil,
		}
		// moving next
		curHead.Next = &nextNode
		curHead = curHead.Next
		l2 = l2.Next
		carry = curSum / 10
	}

	if carry != 0 {
		nextNode := ListNode{
			Val:  carry,
			Next: nil,
		}
		curHead.Next = &nextNode
	}

	return sumHead.Next
}

func addTwoNumbers_Advanced(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, head := 0, l1
	for {
		// calculate
		l1.Val += l2.Val + carry
		carry = l1.Val / 10
		l1.Val = l1.Val % 10
		// exit loop checking
		if l2.Next == nil {
			break
		} else if l1.Next == nil {
			l1.Next = l2.Next
			break
		}
		// move forward
		l1 = l1.Next
		l2 = l2.Next
	}

	// taking care of carry
	for carry != 0 {
		// if no more -> make a new one
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}

		l1.Next.Val += carry
		carry = l1.Next.Val / 10
		l1.Next.Val = l1.Next.Val % 10

		l1 = l1.Next
	}

	return head
}
