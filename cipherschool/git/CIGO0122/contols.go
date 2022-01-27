package main

import "fmt"

/*
1.if-esle statement
     1. if(condition){
		  statements
	  }

	 2. if(condition){
		  statements
	  }else{
		  statements
	  }

	  3if(condition){
		  statements
        }else{
		  statements.}



	switch(data){
	case var1:
		statement
	case var2:
		statament
	case var3:
		statement
	default:
		statement
	}
*/

func main() {
	/*fmt.Println("enter a number:")
	var input int
	fmt.Scan(&input)

	if input%2==0{
		fmt.Println("it is even")
	}else{
		fmt.Println("odd babay")
	}
	*/

	/*x := 10000
	if x = 10; 100%2 == 0 {
		fmt.Println("hey you are even")
		x = 400
	} else {
		fmt.Println("hey you are odd")
	}
	fmt.Println(x)
	*/

	data := 100
	switch data {
	case 1:
		data = 9
		fmt.Println(data)
	case 100:
		data = 200
		fmt.Println(data)
		break
	case 3:
		data = 30
		fmt.Println(data)
		fallthrough //execute the next case also
	default:
		data = 3456
		fmt.Println(data)
	}

}
