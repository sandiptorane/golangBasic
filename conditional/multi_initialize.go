package main

import "fmt"

func main(){
	for i,num := 1,10; i<=10 && num<19; i,num =i+1,num+1{
		fmt.Println(i,"*",num,"=",i*num)
	}
}
