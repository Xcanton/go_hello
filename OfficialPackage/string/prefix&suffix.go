package main

/*
golang 默认解码为 UTF-8， 该编码的优秀特性可以导致许多字符串操作无需解码
下列函数摘自 string包
*/

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[len(s)-len(prefix):] == prefix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
