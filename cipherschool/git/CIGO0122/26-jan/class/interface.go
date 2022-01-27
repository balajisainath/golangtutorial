package main

import (
	"fmt"
	"log"
)

type Integer int

func (i Integer) print() {
	log.Println(i)

}
func (i Integer) calculateSalary() int {
	return int(i)

}

func main() {
	j := Integer(10009)
	j.print()

	//pj := permanentJob{

	//	basicpay: 10,
	//}
	//var sc salaryCalculator
	//sc = pj

	//total := sc.calculateSalary()
	//fmt.Println(total)

	pj := permanentJob{
		basicpay: 10,
	}
	cj := contarctJob{
		basicpay: 20,
	}
	fj := freelanceJob{
		basicpay: 1000,
	}
	i := Integer(1222)

	sc := []salaryCalculator{pj, cj, fj}
	totalIncome(sc)

}
func totalIncome(sc []salaryCalculator) {
	total := 0
	for _, v := range sc {
		total = total + v.calculateSalary()

	}
	fmt.Println(total)
}

//interface- is just a data type- (abstraact data type)
//interface have only methods signature,we can also can have empty interface also
/*
type interface_name interface{
	method-1
	method-2
}
*/

type salaryCalculator interface {
	calculateSalary() int
}
type contarctJob struct {
	basicpay int
}
type freelanceJob struct {
	basicpay int
}

type permanentJob struct {
	basicpay int
}

func (p permanentJob) calculateSalary() int {
	return p.basicpay
}
func (p contarctJob) calculateSalary() int {
	return p.basicpay
}
func (p freelanceJob) calculateSalary() int {
	return p.basicpay
}

//var sc salaryCalculator salaryCalculator
//sc salaryCalculator can hold the data of multiple data types which comes with a certain condtiton
//sc should hold the value of type permanent job and contractJob
//the condition is that types(i.e,permanentJob and contractJob) must implement all the methods of the interface
