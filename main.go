package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Hello world")

	// sum a and b variable
	var a, b int
	a = 1
	b = 2
	fmt.Println("sum: ", a+b)

	// regex
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
}
