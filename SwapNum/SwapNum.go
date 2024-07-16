package main

import "fmt"

func main() {

	fmt.Println("Enter Your Numbers ")
	var n1, n2 int
	temp := 0
	fmt.Scanf("%d %d", &n1, &n2)
	fmt.Println("Numbers Before swapping ", n1, n2)

	temp = n1
	n1 = n2
	n2 = temp

	fmt.Println("Numbers After swapping ", n1, n2)
}
