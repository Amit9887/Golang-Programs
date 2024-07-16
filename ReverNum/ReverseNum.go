package main

import "fmt"

func reverseNum(n int64) int64 {

	var temp, reverse int64 = 0, 0
	for n > 0 {
		temp = n % 10
		reverse = reverse*10 + temp
		n = n / 10

	}
	return reverse
}

func main() {
	fmt.Println("Enter Your Number")
	var num int64
	fmt.Scanln(&num)

	result := reverseNum(num)
	fmt.Println(result)
}
