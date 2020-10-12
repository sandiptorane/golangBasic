package main

import "fmt"

func rectangle(length,width float32)(area float32, perimeter float32){
	 area = length * width
	perimeter = 2*(length+width)
	return  //no explicit return ,it auto return the value

}

func main(){
	var length float32= 12.5
	var width float32= 8.6
	area,perimeter := rectangle(length,width)
	fmt.Println("Area=",area,"\nPerimeter=",perimeter)	

}
