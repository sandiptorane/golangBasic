//This program will print the sum of the squares and cubes of the individual digits of a number.
package main

import (
	"fmt"
)

func squaresSum(number int,squarech chan int){
	sum :=0
	for number !=0 {
		digit := number % 10
		sum = sum + (digit * digit)
		number /=10
	}
	squarech <- sum
}

func cubesSum(number int,cubech chan int){
	sum := 0
	for number !=0 {
		digit := number % 10
		sum = sum + (digit * digit*digit)
		number /=10
	}
	cubech <- sum
}

func main(){
	num := 234
	squarech := make(chan int)
	cubech := make(chan int)
	go squaresSum(num,squarech)
	go cubesSum(num,cubech)
	squares,cubes := <-squarech,<-cubech //receive data from squarech and cubech
	fmt.Println("sum =",squares+cubes)
}


