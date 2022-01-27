package main

import "fmt"

//main function will be automatically called whwn you start the application
func main() {
	var data int
	data = 10
	Data1 := 100 //declare the var and assign to it dtype automatic

	//data1=false//wrong assignment value type
	data = 1000 //in golang if variable is not used that created is error
	fmt.Println(data)
	fmt.Println(Data1)
	//golang is case sensitive

	/*data type
	  1.primitive
	  int,float64,string,bool,complex,byte

	  2.Non-primitive
	   struct,map,chan,interface,array,slice(dynamic size array),root,reflect,pointer

	*/
	name := "doge"
	fmt.Println(name)

}
