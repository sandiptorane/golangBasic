package main

import (  
    "fmt"
)

func main() {

    switch num := 75; { 
    case num < 50:
	if  num < 0 {
	   break
	}
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is lesser than 200", num)
    }

}
