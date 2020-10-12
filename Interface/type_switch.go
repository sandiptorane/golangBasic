package main

import "fmt"

func findType(i interface{}){
	switch i.(type) {
	case int:
		fmt.Printf("type %T value %d\n",i,i)
	case string:
		fmt.Printf("type %T value %s\n",i,i)
	default:
		fmt.Println("unknown type")
	}
}
func main(){
	findType(55)
	findType("hello world")
	findType(65.85)
}