package main

import(
	"fmt"
	"errors"
	"math"
)

func CalculateArea(radius float64)(float64,error){
	if radius<0{
		return 0,errors.New("Area calculation failed,radius should non-negative")
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
