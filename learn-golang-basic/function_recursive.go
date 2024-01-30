package main

import "fmt"

func factorialRecursive(number int) int {
	if number == 10 {
		return 10
	} 
	return number * factorialRecursive(number+1)
}

func main() {
	fmt.Println(factorialRecursive(1))
}