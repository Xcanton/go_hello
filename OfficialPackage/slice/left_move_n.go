package slice

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func left_move_n(s []int, n int) []int {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)

	return s
}
