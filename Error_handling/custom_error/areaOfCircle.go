package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err string
	radius float64
}

func (e *areaError)Error() string{
	return fmt.Sprintf("radius %0.2f : %s",e.radius,e.err)
}

func calculateArea(radius float64)(float64,error){
	if radius < 0{
		return 0,&areaError{"radius is negative",radius}  //return address of areaError with error msg
	}
	return math.Pi*radius*radius,nil
}

func main(){
	radius := -52.5
	area ,err := calculateArea(radius)
	if err!=nil{
		if err,ok := err.(*areaError);ok{
			fmt.Printf("radius %0.2f is less than zero",err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("Area =",area)
}

