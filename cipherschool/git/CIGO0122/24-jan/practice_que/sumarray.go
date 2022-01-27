package main

import "fmt"

func main() {
	var m, n int
	var arrA [10][10]int
	var arrB [10][10]int
	fmt.Println("enter the rows")
	fmt.Scanln(&m)
	fmt.Println("enter the columns")
	fmt.Scanln(&n)
	fmt.Println("enter the array 1")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanln(&arrA[i][j])
		}
	}
	fmt.Println("enter the array 2")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanln(&arrB[i][j])
		}
	}
	fmt.Println("sum of the array is:")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println(arrA[i][j]+arrB[i][j], "t")
		}
		fmt.Println()
	}

}
