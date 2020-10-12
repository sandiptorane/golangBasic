package main

import "fmt"

func compare(s1 string,s2 string){
	if(s1==s2){
		fmt.Println(s1,"and",s2,"are equal")
		return
	}else{
		fmt.Println(s1,"and",s2,"are not equal")
	}
}

func main(){
	var str1 string = "laptop"
	str2 := "computer"
	compare(str1,str2)
	str1 = "Desktop"
	str2 = "Desktop"
	compare(str1,str2)
}
