package main

import "fmt"


func main(){
	var a = []int{1,2,3}
	// 开头添加元素
	fmt.Println(a)
	a = append([]int{0}, a...)
	fmt.Println(a)
	a = append([]int{-3,-2,-1}, a...)
	fmt.Println(a)
}
