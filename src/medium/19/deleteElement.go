package main

import "fmt"


type ListNode struct {
	 Val int
	 Next *ListNode
}

//此题目的难点在于边界条件，如需要删除的倒数第n个元素为头结点。
//这种情况需要单独考虑（带头链表）

//通过设置两个相距为n的结点，这样就不需要记录链表的长度
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var fast, slow *ListNode
	fast = head
	slow = head
	//第一趟遍历到正数第n个值
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	//如果要删除的恰好是头结点
	if fast == nil {
		head = head.Next
		return head
	}
	//此时加入一个 slow 值进行第二趟遍历，slow和fast的距离恰好是n，当fast到达倒数第一个元素时，slow恰好在倒数第n-1个值上
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//将n位置的Next值赋给n-1位置的
	slow.Next = slow.Next.Next
	return head
}

//传统方式
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//链表为空
	if head==nil||n<=0{
		return nil
	}
	p:=head
	p3:=head
	length:=0
	//遍历链表，获取长度
	for p!=nil {
		p=p.Next
		length++
	}

	//对于需要删除头结点的情况特殊处理
	if length==n{
		head=head.Next
		return head
	}
	//修改lenth-n-1位置的值的Next值
	for i:=0;i<length-n-1;i++ {
		p3=p3.Next
	}
	p3.Next=p3.Next.Next
	return head
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

	removeNthFromEnd(listnode1,2)
	fmt.Println("------remove value-------")
	head=listnode1
	for head!=nil{
		fmt.Println(*head)
		head=head.Next
	}

}
