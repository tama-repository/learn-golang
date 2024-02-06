package main

import "fmt"

func main() {
	multiply := multiplyAll
	numList := []int{10, 10, 10, 10}

	fmt.Println(multiply(numList...))
}
