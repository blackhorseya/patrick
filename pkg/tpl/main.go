package tpl

// MainTemplate define a main file
func MainTemplate() []byte {
	return []byte(`package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}
`)
}
