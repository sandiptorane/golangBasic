package main

import(
	"fmt"
	"math"
)

func CalculateArea(radius float64)(float64,error){
	if radius<0{
		return 0,fmt.Errorf("Area calculation failed,radius %0.2f is negative",radius)
	}
	return math.Pi*radius*radius,nil
}

func main(){
	radius := -2.5
	area ,err := CalculateArea(radius)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Area =",area)
}

