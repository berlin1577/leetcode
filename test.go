package main

import "fmt"

func main() {

	s := "jetbrains"
	for i, j := range s {
		fmt.Println(i)
		fmt.Println(j)
	}
	//如果for range 循环结构只有一个返回值，那么就是建

	fmt.Println("this a test")

	fmt.Println("test for debugger")
	//a := 5
	//t := a
	//println(t)
	y := []int{5, 6, 7, 8}
	for i := range y {
		fmt.Println(i)
	}
	for i, h := range y {
		fmt.Println(i)
		fmt.Println(h)
	}
}
