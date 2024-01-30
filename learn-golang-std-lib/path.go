package main

import (
	"fmt"
	"path"
)


func main() {

	pathUrl := "/learn-golang-basic/array.go"

	fmt.Println(path.Base(pathUrl))
	fmt.Println(path.Dir(pathUrl))
	fmt.Println(path.Ext(pathUrl))
	fmt.Println(path.Join("/", "learn-golang-basic", "array.go"))

}
