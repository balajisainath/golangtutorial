package main

import (
	"fmt"
)

func add(num ...int) int { //here num...int means multiple arguments of same data type
	sum := 0
	for value := range num {
		sum += value
	}
	return sum
}
func main() {

	fmt.Println("sum of 1,2,3,4,5=", add(1, 2, 3, 4, 5))
	fmt.Println("sum of 123456789=", add(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
