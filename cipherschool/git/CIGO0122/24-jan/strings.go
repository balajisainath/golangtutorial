package main

import (
	"fmt"
	"strings"
)

//string is consideredas a array of bytes in go
func main() {
	st := "gneojirfnfjsdjzd"
	for _, value := range st {
		fmt.Println(string(value))
	}

	//strings.Contains(st,"a")

	//rest1:=strings.TrimSpace("      tsss     ")//removes unwanted spaces
	res1 := strings.TrimSuffix("tesstarungo", "gotarun")
	fmt.Println(res1)

}
