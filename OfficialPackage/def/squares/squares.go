package main

import "fmt"

func squares() func() int {
	var x int
	// 返回匿名函数，匿名函数默认执行并获取返回值
	// 匿名函数可以获取整个词法环境，即里层函数可以使用外层函数中的变量
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
}
