package main

import "fmt"

func main() {
	fmt.Println("Enter your Array size")
	var size int
	fmt.Scanln(&size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}

	var search int
	fmt.Scanln(&search)

	for index, value := range arr {
		if value == search {
			fmt.Println(search, "element has found at ", index)
		}
	}
}
