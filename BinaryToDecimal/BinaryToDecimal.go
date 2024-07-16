package main

import (
	"fmt"
	"math"
)

func BinaryToDec(n int) int {
	var i int = 0
	var dec int = 0
	for n > 0 {
		rem := n % 10
		n = n / 10
		dec = dec + rem*int(math.Pow(2, float64(i)))
		i += 1
	}
	return dec
}

func main() {
	fmt.Println("Enter Your Number")
	var num int
	fmt.Scanln(&num)
	fmt.Println(BinaryToDec(num))
}
