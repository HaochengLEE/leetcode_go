package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//此题只需要递归两颗树进行比较
func isSameTree(p *TreeNode, q *TreeNode) bool {
	//边界条件
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		//前序遍历
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
func main() {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node1.Right = &node2
	node1.Left = &node3
	node2.Left = &node4
	flag := isSameTree(&node1, &node2)
	fmt.Print(flag)

}
