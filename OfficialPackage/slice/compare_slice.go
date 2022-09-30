package slice

import "fmt"

// slice 唯一能做比较的就是 nil，
// 因为slice的默认值是 nil
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	var s []int // len(s) == 0, s == nil
	s = nil     // len(s) == 0, s == nil
	// s = []int{nil} // len(s) == 0, s == nil 新版本报错
	s = []int{} // len(s) == 0, s != nil
	fmt.Println(cap(fmt.Scan()))
}
