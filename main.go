package main

import "fmt"


func main(){
	var ip *int
	var fp *float32

	var a int= 20
	ip = &a
	fmt.Println(ip);
	fmt.Println(fp);
}
