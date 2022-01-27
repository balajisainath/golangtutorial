package main

import "fmt"

type pill int

const (
	placebo       pill = iota //iota=0,1,2,3,4,5
	Aspirin                   //1
	Tbuprpotein               //2
	Paracetemol               //3
	Acetaminophen             //4
)

/*
var(
	name string
	age intt
)
*/
//go get golang.org/x/tools/cmd/stringer
//stringer -type=Pill

func main() {
	fmt.Println(Paracetemol)
}

//gives output 3 shows paractemol
