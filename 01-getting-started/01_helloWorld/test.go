package main

import "fmt"

var x string = "Hello world"

func main(){
	fmt.Println(x)
	x = "not hello world"
	fmt.Println(x)
}