package main

import "fmt"


type struct_variable_type struct {
	a int
	b int
	c string
}

func main(){
	test := struct_variable_type {1,1,"set"}
	fmt.Println(test.a)
}
