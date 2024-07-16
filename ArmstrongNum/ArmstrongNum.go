package main

import (
	"fmt"
	"math"
)

func checkArmStrong(num int) bool {
	var order = 0
	var temp = num

	for temp != 0 {
		temp /= 10
		order++
	}

	temp = num
	var sum = 0

	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(order)))
		temp /= 10
	}

	return sum == num
}

func main() {
	fmt.Println("Enter your low and high range:")
	var low, high int
	fmt.Scanln(&low, &high)

	var armstrongNumbers []int

	for num := low; num <= high; num++ {
		if checkArmStrong(num) {
			armstrongNumbers = append(armstrongNumbers, num)
		}
	}

	if len(armstrongNumbers) == 0 {
		fmt.Println("The given range does not have any Armstrong numbers.")
	} else {
		fmt.Println("Armstrong numbers within the given range are:")
		for _, armstrong := range armstrongNumbers {
			fmt.Println(armstrong)
		}
	}
}
