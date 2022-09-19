package main

import "fmt"

func main() {
	p := new(int)  // *int类型的p， 指向未命名的int变量
	fmt.PrintLn(*p)  // 输出0
	*p = 2  // 把未命名的int设置为2
	fmt.PrintLn(*p)  // 输出2
}