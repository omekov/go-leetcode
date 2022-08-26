package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// list := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 6,
	// 			Next: &ListNode{
	// 				Val: 3,
	// 				Next: &ListNode{
	// 					Val: 4,
	// 					Next: &ListNode{
	// 						Val: 5,
	// 						Next: &ListNode{
	// 							Val: 6,
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	list := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 7,
				},
			},
		},
	}

	fmt.Printf("%#v\n", *removeElements(list, 7))
}

func Reverse(now *ListNode) *ListNode {
	var cur = now
	var prev *ListNode

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := ListNode{Next: head}
	cur := Reverse(head)
	prev := &dummy

	for cur.Next != nil {

		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}

		cur = cur.Next
	}
	return Reverse(dummy.Next)
}
