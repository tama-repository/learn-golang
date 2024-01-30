package main

import (
	"fmt"
	"learn-golang-basic/helper"
) 


func getHello(name string, filterFunc helper.Filter) string {
	result := filterFunc(name)

	return "Hello " + result
}

func main() {
	fmt.Println(getHello("Trirahmanto", func(name string) string  {
		if name == "Bego" {
			return "..."
		}

		return name
	}))

	spamFunc := func(name string) string  {
		if name == "Bego" {
			return "..."
		}

		return name
	}

	fmt.Println(getHello("Bego", spamFunc))
}