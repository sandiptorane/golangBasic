package main

import "fmt"

func calculate(price,no int) int{
	totalPrice := price*no
	return totalPrice
}

func main(){
   
 	book_price ,no := 100,5
 	total_price := calculate(book_price,no)
	fmt.Println("Total price=",total_price)
}
