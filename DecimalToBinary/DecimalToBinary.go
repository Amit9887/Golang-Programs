package main

import (
	"fmt"
	"strconv"
)

func DecimalToBinary(num int) string {
	if num == 0 {
		return "0"
	}
	var bin string
	for num > 0 {
		rem := num % 2
		num = num / 2
		bin = strconv.Itoa(rem) + bin
	}
	return bin
}

func main() {
	fmt.Println("Enter your Decimal number")
	var dec int
	fmt.Scanln(&dec)
	fmt.Println(DecimalToBinary(dec))
}
