package main
import "fmt"

func main(){
	a,b:=15.20,10.57
	fmt.Printf("type of a is %T and  b is %T",a,b)

	sum:=a+b
	diff:=a-b
	fmt.Println("sum=",sum,"diff=",diff)
}
