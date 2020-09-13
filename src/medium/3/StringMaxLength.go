package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	//判断是不是空串
	if len(s)==0 {
		return 0
	}
	//定义左右边界
	max,left,right:=0,0,0
	//用bool型数组存储结果，为ture即已有该字符
	var maps [256] bool
	for left<len(s){

		if maps[s[right]]{
			maps[s[left]]=false
			//若已有该字符，左边界右移
			left++
		}else {
			//若没有该字符，将该字符位置变为ture，右边界右移
			maps[s[right]]=true
			right++
		}
		if max<right-left{
			max=right-left
		}
		if left+max>=len(s)||right>len(s){
			break
		}
	}
	return max
}

func main() {

	fmt.Print(lengthOfLongestSubstring("absbsb"))
}
