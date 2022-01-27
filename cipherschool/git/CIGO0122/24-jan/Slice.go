package main

import "fmt"

//var arr [4]int

func main() {
	/*arr[0] = 1 //arr index start with zero
	arr[1] = 3
	arr[2] = 4

	fmt.Println(arr) //here arr[4]=0 cze if not assigned it is zero default
	*/

	/*arr := [...]int{1, 2, 3, 4} //here this ... takes the the n val arr automatic can declare any variables b4 assogniing size
	arr1 := [2][2]int{{1, 2}, {3, 4}}
	myarray := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	//iterate array using for loop

	for x := 0; x < len(myarray); x++ {
		fmt.Println(myarray[x])
	}
	fmt.Println(myarray)
	fmt.Println(arr1)
	fmt.Println(arr)
	*/

	/*
		//printing array with index value
		myarray := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
		for index, value := range myarray {
			fmt.Println("element at", index, ":", value)
		}
	*/

	/*myarray1 := [2][2]int{{1, 2}, {4, 3}}
	for index1, value1 := range myarray1 {
		for index2, value2 := range value1 {
			fmt.Println("element at ", index1, " ", index2, "is", value2)
		}
	}
	*/

	//Slice

	s := make([]int, 5, 15)
	//make ([]data_type,length,capacity) make syntax
	//var slice []int

	for index := range s {
		s[index] = 5
	}
	for i := 0; i < 6; i++ {
		s = append(s, 20)

	}
	s = append(s, 25)
	s = append(s, 23)
	s = append(s, 26)
	fmt.Println(s)

	// making the slice of the array

	arr := []int{0, 3, 2, 5, 6, 4}
	arr1 := arr[0:5] //from arr[0] to arr[4] excluding arr[5]on wards
	arr1[0] = 10     // it will change the original slice as well

}
