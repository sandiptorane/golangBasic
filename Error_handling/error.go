package main

import (
	"fmt"
	"os"
)

func main(){
	file,err := os.Open("/abc.text")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name(),"opened successfully")

}
