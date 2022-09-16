package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // bufio.NewScanner 是一个扫描器，能够通过调用 Scan() 读取下一个单位（行[默认] or 单词），并清除行尾的换行符。
		counts[input.Text()]++
	}

	// 注意： 忽略  input.Err()  中可能出现的错误

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
