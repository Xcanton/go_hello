package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)

	fmt.Println(y, strconv.Itoa(x))

	// FormatInt 和 FormatUint 可以按不同的进制位格式化数字

	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"

	s := fmt.Sprint("x=%b", x) // "x=1111011"

	// strconv包内的Atoi函数或者ParseInt函数用于解释表示整数的字符串
	// ，而ParseUint用于无符号数

	x_, _ := strconv.Atoi("123")
	y, err := strconv.ParseInt("+123", 10, 64)
}
