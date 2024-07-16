package main

import "fmt"

func FindPrimeFactor(number int) {
	for i := 2; i < number; i++ {
		for number%i == 0 {
			number = number / i
			fmt.Println(i)
		}
	}
	if number != 1 {
		fmt.Println(number)
	}
}

func main() {
	fmt.Println("Enter the Number")
	var num int
	fmt.Scanln(&num)
	FindPrimeFactor(num)
}
