package main

import (
	"LeetCode/basic_algorithm"
	"fmt"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	ans := basic_algorithm.MinWindow(s, t)
	fmt.Println(ans)
	//fmt.Println("fdfs")
	//basic_algorithm.Test()
}
