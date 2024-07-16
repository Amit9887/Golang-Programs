package main

import "fmt"

func isVowelCons(s *string) int {
	var c, v = 0, 0

	for _, char := range *s {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			v++
		case ' ', ',', '!':
			continue
		default:
			c++
		}
	}
	return c + v
}

func main() {
	fmt.Println("Enter your String")
	var str string
	fmt.Scanln(&str)

	result := isVowelCons(&str)
	fmt.Println(result)
}
