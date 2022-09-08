package main

import "fmt"


func main(){
	var valance = [5]float32{100.0, 2.0, 3.7, 7.0, 50.0}
	fmt.Println(valance)
	balance := [5]float32{110.0,2.0,50.0}
	fmt.Println(balance)
	balance2:=[...]float32{100.0, 2.0,3.7,7.0,50.0}
	fmt.Println(balance2)
	var valance1 = [5]float32{1:2.0, 3:7.0}
	fmt.Println(valance1)
}
