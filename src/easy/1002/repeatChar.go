package main

import (
	"fmt"
	"math"
)

//1.需要统计每个单词中都有的字母 2.需要统计单词中重复的字母


func commonChars(A []string) []string {

	//两个数组，一个统计每个单词的字母，一个统计所有单词重复的字母
	arr:=[26]int{}
	for i:=range arr{
		arr[i]=math.MaxUint16
	}
	num:=[26]int{}
	for _,word:=range A{
		//统计单词的字母
		for _,char:=range []byte(word){
			num[char-'a']++
		}
		//将 num 更新到 arr 中
		for i:=0;i<26;i++ {
			//如果 arr[i] 中没有值，则更新到 arr
			if arr[i]>num[i] {
				arr[i]=num[i]
			}
		}
		//清空 num
		for i:= range num{
			num[i]=0
		}
	}
	result:=make([]string,0)
	//遍历arr
	for i:=0;i<26;i++ {
		for j:=0;j<arr[i];j++ {
			result=append(result,string(i+'a'))
		}
	}
	return result
}
func main() {
	words:=[]string{"giao","oka","jiao"}
	newwords:=commonChars(words)
	for _,i:=range newwords{
		fmt.Printf(i)

	}


}
