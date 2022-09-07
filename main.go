package main

import "fmt"

func Add(a, b int) int {
	return a+b
}

func main(){
	fmt.Println(Add(2,3))
	var Sub = func(a,b int) int {return a-b}
	fmt.Println(Sub)
	fmt.Println(Sub(3,1))
}
