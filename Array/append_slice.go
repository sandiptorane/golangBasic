package main

import (
    "fmt"
)

func main(){
	var names []string //zero value of slice is nil
	if names == nil {
		fmt.Println("slice nil going to append")
		names = append(names,"Sandip","Dinesh","Yash")
		fmt.Println("names=",names)
	}
	names= append(names,"Dhiraj")
	fmt.Println("names after second append=",names)
      
       names2 := []string{"Nagesh","Uday","Avdhut"}
	names= append(names,names2...)
	fmt.Println("names after  appending another slice to names=",names)

}
