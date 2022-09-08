package main

import "fmt"


func GetData() (int, int) {
	return 10, 20
}

func main(){
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a,b)
}
