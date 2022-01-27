package main

import "fmt"

func main() {
	var num int
	println("enter the positive integer")
	fmt.Scanln(&num)
	result := palindrome(num)
	fmt.Println("result is", result)
}
func palindrome(num int) bool {
	var reminder, temp int
	var reverse int = 0
	temp = num

	for {
		reminder = num % 10
		reverse = reverse*10 + reminder
		num /= 10
		if num == 0 {
			break
		}
	}
	if temp == reverse {
		return true
	} else {
		return false
	}

}
