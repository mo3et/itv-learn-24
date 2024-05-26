package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next // 先保存下个节点的指针
		cur.Next = pre   // 将cur的Next 设置为 pre (反转，当前节点指向前节点)
		pre = cur        // 移动 pre 指针到当前节点cur
		cur = next       // 移动 cur 到下个节点next
	}
	return cur
}
