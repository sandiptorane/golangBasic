package main

import(
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup){
	fmt.Println("gouroutine started",i)
	time.Sleep(2*time.Second)
	fmt.Println("gouroutine ended",i)
	wg.Done()       //waitGroups counter decreamented
}

func main(){
	num := 4
	var wg sync.WaitGroup
	for i:=0;i<num;i++{
		wg.Add(1) //waitGroup counter  increamented by  value passed to Add
		go process(i,&wg)  //compulsory pass address of waitgroup
	}
	wg.Wait() //blocks the goroutine until counter becomes zero i.e all until executes
	fmt.Println("All go routines finished")
}
