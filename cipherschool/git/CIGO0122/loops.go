package main

import "fmt"

/**/

func main() {

	/*	for i :=0;i<3;i++{
			for j:=0;i<3;j++{
			fmt.Println(j)
			}
		}

		for true{
			fmt.Println("infinite executtion")}
	*/
	/*	i:=0
		for i<3{
				fmt.Println(i)
				i++
			}*/

	i := 0 //do while by for loop
	for {
		fmt.Println(i)
		i++

		if i >= 3 {
			break
		}
	}
	/*if data:=10;data<=10{
		fmt.Print("less")
	}else{
		fmt.Print("more")
	}
	*/

}
