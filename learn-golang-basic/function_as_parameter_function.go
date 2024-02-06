package main

import "fmt"

type Filter func(string) string

func spamFilter(name string) string {
	if name == "Bego" || name == "Tolol" {
		return "<censored>"
	}
	return name
}

func sayHello(name string, filterFunc Filter) string {
	nameFiltered := filterFunc(name)
	return "Hello " + nameFiltered
}

func main() {
	fmt.Println(sayHello("Bego", spamFilter))
	fmt.Println(sayHello("Hutama", spamFilter))
}
