package main

import (
	"fmt"
	"os"
)

/*type error interface {
	Error() string
}

type PathError struct {
	Op string
	Path string
	Err error
}
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}*/

func main(){
	file,err := os.Open("test.txt")
	if err!=nil{
		if pErr,ok := err.(*os.PathError);ok{
			fmt.Println("failed to open file at ",pErr.Path)
			return
		}
		fmt.Println("generic error",err)
	}
	fmt.Println(file.Name(),"file opened successfully")
}