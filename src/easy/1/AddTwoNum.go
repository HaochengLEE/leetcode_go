package main

import "fmt"

//暴力破解，两个for循环，时间复杂度O(n^2)
func twoSum1(nums []int,target int) []int{
	for i,v:=range nums{
		for k:=i+1;k<len(nums);k++{
			if target-v==nums[k]{
				return []int{i,k}
			}
		}
	}
	return nil

}

//使用Map存储数字，判断Map中是否有对应的值
func twoSum2(nums []int, target int) []int {
	m:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		a:=target-nums[i]
		if _ , ok:=m[a];ok{
			return []int{m[a],i}
		}
		m[nums[i]]=i
	}
	return nil
}

func main() {
	array:=[]int {1,2,3,4,5,6}
	fmt.Print(twoSum2(array,5))

}
