package main

import (
	"fmt"
)

/*defining map syntax

map[key_Type]value_Type{}
var my map[int] string

*/

func main() {
	studentclass := map[int]string{
		1: "aditya",
		2: "shranya",
		3: "devanshu",
	}
	studentclass[15] = "vishnu" //adding keys values
	studentclass[17] = "praveen"
	/*	for roll_no, name := range studentclass {
			fmt.Println("Roll no:", roll_no, "name:", name)
		}
	*/
	student, ok := studentclass[5] //to check if that key is valid or there
	if ok {
		fmt.Println(student)

	} else {
		fmt.Println("key invalid")
	}
	delete(studentclass, 15) //this used to delete key
	for roll_no, name := range studentclass {
		fmt.Println("roll no:", roll_no, "name", name)
	}
	new_map := studentclass
	new_map[15] = "rohan" //newmap of studentclass
	//sort.Ints()

	studentclass2 := map[int]string{
		21: "gitanshu",
		22: "sara",
		35: "tara",
	}
	studentclass[15] = "vishnu"
	studentclass[17] = "tarun"
	for key, value := range studentclass2 {
		studentclass[key] = value
		delete(studentclass[21]) = "gitanshu"
	}
}
