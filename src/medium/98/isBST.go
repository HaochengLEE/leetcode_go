package main

import (
	"fmt"
)

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	flag:=true
	result:=isBST(root,flag)
	//限定最大最小值
	//result:=isBST2(root,math.Inf(-1),math.Inf(1))
	return result

}
func isBST2(root *TreeNode,min,max float64) bool  {
	//到底，返回true
	if root==nil {
		return true
	}
	//将当前结点的 Val 和左右结点比较
	v:=float64(root.Val)
	return v<max&&v>min&&isBST2(root.Left,min,v)&&isBST2(root.Right,v,max)
}
func isBST(root *TreeNode,flag bool) bool {
	if root==nil{
		return true
	}
	//与左右值相比较
	left:=*root.Left
	right:=*root.Right
	if left.Val>root.Val||right.Val<root.Val {
		return false

	}
	return isBST(&left,flag)&&isBST(&right,flag)

}
func main() {
	node1:=TreeNode{Val: 1}
	node2:=TreeNode{Val: 2}
	node3:=TreeNode{Val: 3}
	node4:=TreeNode{Val: 4}
	node1.Right=&node2
	node1.Left=&node3
	node2.Left=&node4
	flag:=isValidBST(&node1)
	fmt.Print(flag)
	
}
