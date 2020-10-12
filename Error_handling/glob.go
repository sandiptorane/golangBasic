package main

import (
	"errors"
	"fmt"
	"path/filepath"
)

var ErrBadPattern = errors.New("syntax error in pattern")

func main(){
	file,err := filepath.Glob("[")
	if err !=nil{
		if err==filepath.ErrBadPattern{
			fmt.Println("Bad pattren error",err)
			return
		}
		fmt.Println("Generic error",err)
		return
	}
	fmt.Println("succesfully opened matching file",file)
}
