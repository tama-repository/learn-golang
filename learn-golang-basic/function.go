package main

import "fmt"


func greeting() {
	fmt.Println("Hello Hutama Trirahmanto")
}

func sumValue(num1 int, num2 int) int {
	result := num1 + num2
	return result
}
func main() {

	greeting()
	fmt.Println(sumValue(5, 10))

}