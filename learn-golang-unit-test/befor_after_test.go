package main

import (
	"fmt"
	"testing"
)


func TestMain(m *testing.M) {
	fmt.Println("Before test run")

	m.Run()

	fmt.Println("After test run")
}
