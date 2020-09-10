package main

import "fmt"

func combine(n int,k int)  {
	fmt.Println("{")
	for key:=n;key>0;key--{
		fmt.Printf("[")
		for v:=k;v>0;v--{
			fmt.Printf("%d",v)
			if v>1{
				fmt.Printf(",")
			}

		}
		fmt.Println("]")
	}
	fmt.Println("}")

}
func main()  {
	combine(4,2)

}
