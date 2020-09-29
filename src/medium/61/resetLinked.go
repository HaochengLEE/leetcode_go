package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
func rotateRight1(head *ListNode, k int) *ListNode {
	//边界条件
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	//声明一个新的头结点
	newHead := &ListNode{Val: 0, Next: head}
	len := 0
	cur := newHead
	//遍历链表获取length
	for cur.Next != nil {
		len++
		cur = cur.Next
	}
	//边界条件，不需要挪动的情况
	if (k % len) == 0 {
		return head
	}
	//重置cur
	cur.Next = head
	cur = newHead
	//遍历到需要挪动的元素前的元素
	for i := len - k%len; i > 0; i-- {
		cur = cur.Next
	}
	//res存储要挪动元素的地址
	res := &ListNode{Val: 0, Next: cur.Next}
	cur.Next = nil
	return res.Next
}
func rotateRight(head *ListNode, k int) *ListNode {
	//去除边界条件
	if head==nil||head.Next==nil||k==0{
		return head
	}
	//链表长度
	length:=0
	//哨兵结点
	newhead:=&ListNode{0,head}

	p1:=newhead
	for p1.Next!=nil{
		length++
		p1=p1.Next
	}

	p1.Next=head
	p1=newhead
	//无需挪动的情况
	if k%length==0{
		return head
	}
	//长度对k取余求出需要挪动的个数
	for i:=length-k%length; i>0 ;i-- {
		p1=p1.Next
	}

	res:=&ListNode{0,p1.Next}
	p1.Next=nil
	return res.Next
}
func main() {
	var listnode1= new(ListNode)
	listnode1.Val=1
	var listnode2 = new(ListNode)
	listnode2.Val=2
	listnode1.Next=listnode2
	var listnode3 = new(ListNode)
	listnode3.Val=3
	listnode2.Next=listnode3
	var listnode4 = new(ListNode)
	listnode4.Val=4
	listnode3.Next=listnode4
	head:=listnode1
	for head!=nil{
		fmt.Println(*head)
		head=head.Next
	}

	fmt.Println("------remove value-------")
	head=rotateRight(listnode1,3)
	for head!=nil{
		fmt.Println(*head)
		head=head.Next
	}


}
