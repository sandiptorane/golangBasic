package main

import "fmt"

func rectangle(length,width float32)(float32,float32){
	 area := length * width
	perimeter := 2*(length+width)
	return area,perimeter

}

func main(){
	var length float32= 12.5
	var width float32= 8.6
	area, _ := rectangle(length,width) //perimeter discarded
	fmt.Println("Area=",area)	

}
