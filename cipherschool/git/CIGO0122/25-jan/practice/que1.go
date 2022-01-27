package main

import "fmt"

//given an array[]of N non-negative integers representing the height of blocks if thw width of each block 1  how much water it can hold in raint season

var arr=[]int{3,0,0,2,0,4}
var maxHeight int
func min(x,y int)int{

	if x<y{
		return x
	}
	return y
}
func volume(h int)int{
	return h
}

func main(){
	maxHeight=min(arr[0],arr[len(arr)-1])
	waterCaptured:=0
	for _,value:=range arr{
		var waterHeight int
		if value>maxHeight{
			value=maxHeight
		}
		if index>0 && index<len(arr-1){
			HeightLeft:=arr[index-1]
			HeightRight:=arr[index+1]
			if HeightLeft>maxHeight && HeightRight >maxHeight{
				waterHeight
			}

		}else
		waterHeight:=maxHeight-value
		waterCaptured+=volume(waterHeight)
	}
	fmt.Println("total water covered=",waterCaptured)
	
}