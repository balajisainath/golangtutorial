package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p person) Details() string {
	return fmt.Sprintf("%v(%v years)", p.Name, p.Age)
}

func main() {

	a := person{"tarun", 27}
	fmt.Println(a)
	//fmt.Println("name=", a.Name, "age=", a.Age)
}
