package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	words := strings.Fields(str)
	reversed := ""
	for i := len(words) - 1; i >= 0; i-- {
		reversed += words[i] + " "
	}
	fmt.Println("Reversed:", reversed)
}
