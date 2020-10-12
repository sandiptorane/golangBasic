package main

import "fmt"

func main(){
	employee := map[string]int{
		"steve": 12000,
		"sam" : 12400,
		"mike" : 15440,
	}
	fmt.Println("map before modification :",employee)
	modifyEmployee := employee  //reference to modifyEmployee like array to slice
	modifyEmployee["sam"]=13550 //sam's salary updated in modifyEmployee
	fmt.Println("map after modification :",employee)
}
