package main
import "fmt"

func main(){
	c1:=complex(6,8)
	c2:= 8 + 27i
	add := c1+c2
	mul := c1*c2
	fmt.Println("sum:",add)
	fmt.Println("product",mul)
}

