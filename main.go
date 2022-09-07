package main

import "fmt"


func main(){
	var a []int
	fmt.Println(a)
	a = append(a, 1)
	fmt.Print(a)
	a = append(a, 1,2,3)
	fmt.Println(a)
	a = append(a, []int{1,2,3}...)
	fmt.Println(a)
}
