package main

import "fmt"

// this file cannot be executed directly
func main() {
	var x int = %d
	const str string = "%s"
	for i := 0; i < x; i++ {
		fmt.Println(str)
	}
}
