package main

import (
	"os"
	"strconv"
	"fmt"
)

// var sliceb []int

func main() {
	var sliceA []int
	//reads from the terminal
	if len(os.Args) > 2 {
		return
	}
	args := os.Args[1]
	for _, arg := range args {
		num, _ := strconv.Atoi(string(arg))
		sliceA = append(sliceA, num)
	}
	fmt.Println(rra(sliceA))
	fmt.Println(ra(sliceA))
	fmt.Println(sa(sliceA))
	

}

//swap top two numbers
func sa(num []int) []int {
	if len(num) > 1 {
		num[0], num[1] = num[1], num[0]
	}
	return num
}
//Move the top number in stack a to the bottom.
func ra(num []int) []int {
if len(num) > 1  {
	num=append(num[1:],num[0])
}
 return num
}
//rra: Move the bottom number in stack a to the top.
func rra(num []int) []int {
if len(num) > 1  {
	num= append(num[len(num)-1:],num[:len(num)-1]...)
}
return num
}
//pa: Take the top number from stack b and put it on top of stack a.
