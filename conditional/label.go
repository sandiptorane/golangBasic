package main

import "fmt"

func main(){
outer:
	for i:=0;i<3;i++ {
		for j:=1;j<4;j++{
			fmt.Println("i=",i,"j=",j)
			if i==j{
			   break outer
			}
		}
	}

}
