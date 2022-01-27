package main

import "fmt"

/*
structure is just a bluprint
idea,it wont consume any of the ram space untill and unless the instance of it is created


structure -instance of structure
like
class-object
//10 different -consume same space

syntax
type struct_name struct{

}
*/

type Iphone13 struct {
	camera float32 //below are called attributes of struccture
	color  string
	ram    int
	rom    int
	price  float32
	size   float32
	old    Iphone12 //nested struct
	//method
	//call func()

}
type Iphone12 struct {
	camera int
}

func (i *Iphone13) call(number string) { //only i Iphone13 gives pass by value
	//only  (i *Iphone13) it will give pass by refrence
	i.camera = 1222223
	fmt.Println(i.camera) //thiz give default val i.e 0.0
	fmt.Println("calling....." + number)
}
func main() {
	//pass by reference
	d := &Iphone13{} //pass by value if d:=Iphone13{}
	d.camera = 10
	d.call("999100")
	fmt.Println(d.camera)
	fmt.Println(d)
	fmt.Println(*d)

	/*	//create instace of iphone
		//empty
		a := Iphone13{}
		a.call("9123456788")
		a.camera = 12.4
		a.color = "black"
		fmt.Println(a)        //if i implement is o/p {12.4,black,0,0} no garbage only default is given if attributes not given valuess
		fmt.Println(a.camera) //this print only the 'a.camera value'
		b := Iphone13{}
		b.call("09999")
		a.call("90345678")

		//non empty instance one method
		//c := Iphone13{
		//		camera: 10,
		//	}
		c := Iphone13{camera: 10,
			ram:   12,
			rom:   12,
			size:  16,
			color: "silver",
			price: 1,
			old: Iphone12{
				camera: 111,
			},
		} //here coma is not needed
		fmt.Println(c)
		fmt.Println(a.old.camera) //printing the camera from iphone 12
	*/
}
