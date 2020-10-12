package main

import "fmt"

func describe(i interface{}){
	fmt.Printf("Type = %T value = %v\n",i,i)
}

func main(){
	i := 55
	describe(i)
	name := "hello world"
	describe(name)
	person := struct{
		firstname string
		lastname string
	}{
		firstname: "Sandip",
		lastname: "Torane",
	}
	describe(person)
}

