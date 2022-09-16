package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("第一项是 执行文件的绝对路径" + os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
