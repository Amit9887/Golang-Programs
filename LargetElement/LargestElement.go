package main

import "fmt"

func main() {
	fmt.Println("enter your array size")
	var size int
	fmt.Scanln(&size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}

	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	fmt.Println(max)
}
