package main

import "fmt"

func main() {
	var x int = %d
	const str string = "%s"
	for i := 0; i < x; i++ {
		fmt.Println(str)
	}
}
