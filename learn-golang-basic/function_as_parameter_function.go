package main

import "fmt"

func spamFilter(name string) string {
	if name == "Bego" {
		return "..."
	} 
	return name
}

func main() {
	fmt.Println(getHello("Bego", spamFilter))
	fmt.Println(getHello("Hutama", spamFilter))
}