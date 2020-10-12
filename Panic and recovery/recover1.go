package main

import "fmt"

func recoverfullname(){ //if panic occurs it will recover
	if r:=recover();r!=nil{
		fmt.Println("recoverd from",r)
	}
}
func fullname(firstname *string,lastname *string){
	defer recoverfullname()
	if firstname==nil{
		panic("runtime error: fisrtname cannot be nil")
	}
	if lastname==nil{
		panic("runtime error: lastname cannot be nil")
	}
	fmt.Println("fullname =",*firstname,*lastname)
	fmt.Println("successfully terminated from fullname")
}

func main(){
	defer fmt.Println("defered call in main")
	firstname := "jack"
	//lastname :="jobs"
	fullname(&firstname,nil)
	fmt.Println("successfully terminted from main")
}

