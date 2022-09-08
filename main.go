package main

import "fmt"


func main(){
	a := []int{1,2,3,4,5,6,7,8}
	fmt.Println(a)
	var b = append(a[:0], a[1:]...)
	fmt.Println(b)
	//var c = a[:copy(a, a[3:])]
	//fmt.Println(c)
}
