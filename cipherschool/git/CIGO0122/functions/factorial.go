package main

import (
	"errors"
	"fmt"
)

func factorial(number int) (fact int, err error) {
	fact = 1
	if number < 0 {
		err = errors.New("enter the positive number")
		return
	} else if number == 0 {
		return

	}
	for number > 0 {
		fact *= number
		number--

	}
	return

}

func main() {
	var number int
	fmt.Println("enter the positive number")
	fmt.Scanln(&number)

	result, err := factorial(number)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("factorial is ofis=", result)
	}

}
