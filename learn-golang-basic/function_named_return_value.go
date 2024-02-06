package main

import "fmt"

func fullNameFunc() (firstName, middleName, lastName string) {
	firstName = "Hutama"
	middleName = "Tri"
	lastName = "Rahmanto"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := fullNameFunc()
	fmt.Println(a, b, c)
}
