package main
type ListNode struct {
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node:=new(ListNode)
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
func main() {
	node1:=new(ListNode)

	
}
