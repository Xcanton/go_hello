package main

import "fmt"


func main(){
	// go中数组地址代表整一个数组，并非c中头节点指针，所以赋值传递数组时通常只使用指针，减小开销
	var a = [...]int{1,2,3}
	var b = &a

	for i,v := range b { fmt.Println(i,v) }
}
