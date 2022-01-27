package main

import (
	"errors"
	"fmt"
)

func palinfrome(word string) (string, error) {

	if word == "" {
		return word, errors.New("the word cannot be empty")
	} else {
		return "hello " + word, nil
	}
}
func main() {
	greetings, err := palinfrome("peter!!")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greetings)
	}

}
