package main

import "fmt"

func main() {
	fmt.Println("Enter your number")
	var n1, n2 int
	fmt.Scanln(&n1, &n2)

	on1 := n1
	on2 := n2

	for n1%n2 != 0 {
		rem := n1 % n2
		n1 = n2
		n2 = rem
	}

	gcd := n2
	lcm := (on1 * on2) / gcd

	fmt.Println(gcd)
	fmt.Println(lcm)

}
