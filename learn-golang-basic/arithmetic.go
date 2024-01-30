package main

import "fmt"


func main() {
num1 := 100
num2 := 350
sum := num1 + num2
red := sum - 50
mul := red * 2
div := mul / 2
mod := div % 2
fmt.Println(sum, red, mul, div, mod)
}