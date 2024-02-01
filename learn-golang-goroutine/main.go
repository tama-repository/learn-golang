package main

import "fmt"



func MultiplyNum(num1, num2 int, channel chan <- int) {
	result := num1 * num1
	channel <- result
}

func DisplayNumber(num int) {
	fmt.Println("Number", num)
}

func main() {
	DisplayNumber(19)
}
