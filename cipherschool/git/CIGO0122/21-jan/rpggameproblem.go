package main

import "fmt"

var id int
var isKnightSleep bool   //decides fast attack
var isPrisonerSleep bool //decides signal prison
var isArcherSleep bool   //decides signal prison
var hasDog bool

//var isAllsleep
func fastattack() bool {
	return isKnightSleep
}
func cansoy() bool {
	return !(isKnightSleep && isArcherSleep && isPrisonerSleep)
}
func canfreeprisoner() bool {
	return ((hasDog && isArcherSleep) || ((isKnightSleep) && isArcherSleep && !(isPrisonerSleep)))

}
func CanSignalPrisoner() bool {
	return (!isPrisonerSleep) && isArcherSleep
}
func main() {
	isArcherSleep, isPrisonerSleep, isKnightSleep, hasDog = true, true, true, true
	fmt.Println("the result is:", fastattack(), cansoy(), canfreeprisoner(), CanSignalPrisoner())
}
