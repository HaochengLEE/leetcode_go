package main

import "fmt"

//暴力，两个for循环，时间复杂度接近n^2
func maxArea(height []int) int {
	var area int
	var h int
	for i:=0;i<len(height);i++{
		for j:=i;j<len(height);j++ {
			if height[i]>height[j]{
				h=height[j]
			}else {
				h=height[i]
			}
			if area<(j-i)*h{

				area=(j-i)*h
			}
		}
	}
	return area
}
//标记左右端点，少一次for循环
func maxArea1(height []int) int {
	max,left,right:=0,0,len(height)-1

	for left<right{
		width:=right-left
		h:=0
		if height[left]<height[right]{
			h=height[left]
			left++
		}else {
			h=height[right]
			right--
		}

		area:=h*width
		if max<area{
			max=area
		}
	}
	return max
}
func main() {
	arr:=[]int {2,3,4,5,18,17,6}
	fmt.Print(maxArea1(arr))
}
