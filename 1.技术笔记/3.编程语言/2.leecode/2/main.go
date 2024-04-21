package main

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {

	//算法思路
	//1.每个节点依次相加
	//1.1 节点的数值（sum）= (l1.val + l2.val + carry(进位值)) % 10
	//1.2 carry (进位值) = (l1.val + l2.val + carry(进位值)) / 10
	//2.每次得到 sum 挂载到结果链表的后面
	//3.如果最后还有一个 carry 则单独将其挂载到链表的最后
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		//获取节点上的值
		n1, n2 := 0, 0
		//获取l1的值
		if l1 != nil {
			n1 = l1.Val
			//切换到下一个节点
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			//切换到下一个节点
			l2 = l2.Next
		}

		//carry = (n1 + n2) / 10
		//获取 两个节点值 + 进位置
		sum := n1 + n2 + carry
		//sum := (n1 + n2 + carry) % 10
		//计算最终节点值和下一次相加的进位值
		sum, carry = sum%10, sum/10
		//如果链表为空，就添加头结点
		if head == nil {
			//添加头结点
			head = &ListNode{Val: sum}
			//tail 指针指向头结点
			tail = head
		} else {
			//设置头结点的next结点
			tail.Next = &ListNode{Val: sum}
			//tail指向新添加的节点，方便下次操作
			tail = tail.Next
		}
	}

	//如果 carry 不为空，就将其作添加在最后一个节点
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
		tail = tail.Next
	}

	return
}
