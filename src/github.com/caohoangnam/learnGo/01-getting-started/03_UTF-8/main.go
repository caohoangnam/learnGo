package main

import "fmt"

func main() {
	for i := 60; i < 120; i++ {
		fmt.Println("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
