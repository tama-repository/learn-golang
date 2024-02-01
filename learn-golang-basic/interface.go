package main

import (
	"fmt"
	"learn-golang-basic/helper"
)

func getUserDetail(user helper.Users) {
	firstName, lastname, address := user.GetUserDetail()
	fmt.Println("FirstName", firstName, "LastName", lastname, "Address", address)
}

func main() {

customerOne := helper.Customer{
	FirstName: "Hutama",
	LastName: "Trirahmanto",
	Address: "Tangerang",
}

	getUserDetail(customerOne)

	// Interface{} / any

var customerTwo any = "Hutama"
customerTwo = 1

fmt.Println(customerTwo)

}
