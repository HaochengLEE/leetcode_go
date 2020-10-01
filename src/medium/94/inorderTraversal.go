package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	}
func inorderTraversal(root *TreeNode) []int {
	//int数组存值
	arr:=[]int{}
	p:=root
	//递归遍历树
	inorder(p,&arr)
	return arr
}
func inorder(p *TreeNode,arr *[]int)  {
	//中序遍历，先左后右
	if p != nil {
		inorder(p.Left, arr)
		*arr = append(*arr, p.Val)
		inorder(p.Right, arr)
	}

}
func main() {
	node1:=TreeNode{Val: 1}
	node2:=TreeNode{Val: 2}
	node3:=TreeNode{Val: 3}
	node4:=TreeNode{Val: 4}
	node1.Right=&node2
	node1.Left=&node3
	node2.Left=&node4
	flag:=inorderTraversal(&node1)
	fmt.Print(flag)
}
