package main

import "fmt"

type AreaError struct{
	err string // descriptiom about err
	length float64
	width float64
}

func (e *AreaError)Error() string{
	return e.err
}
func (e *AreaError)lengthNagative() bool{
	return e.length<0
}
func (e *AreaError)widthNegative() bool{
	return e.width < 0
}

func rectArea(length,width float64)(float64,error){
	err :=""
	if length < 0{
		err += "length is negattive"
	}
	if width<0{
		if err==""{
			err = "width is negative "
		} else{
			err+="width is negative"
		}
	}
	if err!=""{
		return 0,&AreaError{err,length,width}
	}
	return length*width,nil
}

func main(){
	length,width := 25.1,-10.2
	area,err := rectArea(length,width)
	if err,ok := err.(*AreaError);ok{
		if err.lengthNagative(){
			fmt.Println("length",err.length,"is less than zero")
		}
		if err.widthNegative(){
			fmt.Println("width",err.width,"is less than zero")
		}
	}
	fmt.Println("Area of rectangle =",area)
}