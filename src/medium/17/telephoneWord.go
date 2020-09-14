package main

import "fmt"

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res   = []string{}
	final = 0
)
//抄来的解法，DFS
func letterCombinations1(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
	return
}
//自己的解法，
//思路：
//1.数组存储每个数字对应的字母
//2.将为存在的数字放进一个新数组中
//3.组合字母

func letterCombinations(digits string) []string {
	str:=[10]string{"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	flag:=[1000]bool{false}
	data:= []string{}
	result:=[]string{}
	for i:=0;i<len(digits);i++ {
		if flag[digits[i]]!=true{
			data=append(data,str[digits[i]])
			flag[digits[i]]=true
		}

	}
	fmt.Print(data)
	return result
}

func main() {
	letterCombinations("123")
}
