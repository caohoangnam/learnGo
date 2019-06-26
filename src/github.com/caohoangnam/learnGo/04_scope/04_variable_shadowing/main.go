package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	maxx := max(7)
	fmt.Println(maxx) // maxx is now the result, not the function
}

// don't do this; bad coding practice to shadow variables
