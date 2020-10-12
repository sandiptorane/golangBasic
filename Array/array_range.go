package main

import (
	"fmt"
)

func main(){
	a :=[...]float64{64.7,65.3,54.5,64.5,87.3}
	for i:=0; i<len(a);i++{    //looping from 0 to length of array
		fmt.Printf("%dth value =%.2f\n",i,a[i])
	}
	sum := float64(0)
       for _,value := range a{  //range returns index and value both
			sum += value
	}
	fmt.Println("\n sum of array elements=",sum)
}
