package main

import "fmt"

func printarray(arr [3][2]string){
	for _,value1 := range arr{
		for _,value2 := range value1{
			fmt.Printf("%s ",value2)
		}
		fmt.Println()
	}
}

func main(){
	a:= [3][2]string{    //using short hand declaration
		{"lion","tiger"},
		{"cat","dog"},
		{"pigeon","peacock"},
	}
	printarray(a)
       	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] =  "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "Redmi"
	fmt.Printf("\n")
	printarray(b)

}
