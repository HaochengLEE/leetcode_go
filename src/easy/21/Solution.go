package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//创建一个新的链表
	node:=new(ListNode)
	node.Next=new(ListNode)
	//存储链表的头节点
	head:=node


	//var node1 ListNode
	//遍历链表
	//判断值的大小，插入新的链表
	for l1!=nil&&l2!=nil {
		if l1.Val<=l2.Val {
			head.Next=l1
			l1=l1.Next
		}else {
			head.Next=l2
			l2=l2.Next
		}

	}

	if l1==nil {
		head.Next=l2
	}else {
		head.Next=l1
	}

	return node.Next

}
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func main() {
	node1,node2:=new(ListNode),new(ListNode)
	node1.Next=new(ListNode)
	head1:=node1

	for i:=0;i<6;i++{
		node1.Val=i
		node1.Next=new(ListNode)
		fmt.Print(node1.Val)
		node1=node1.Next


	}

	for i:=0;i<6;i++{

		fmt.Print(head1.Next)
	}


	for i:=1;i<6;i++{
		node2.Val=i
		node2.Next=new(ListNode)
		node2=node2.Next
	}

	//fmt.Print(mergeTwoLists1(node1,node2))

	
}
