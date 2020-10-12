package main

import "fmt"

type worker interface {
	work()
}

type person struct{
	name string
}

func (p person)work(){
	fmt.Println(p.name,"is working in company")
}
func describe(w worker){
	fmt.Printf("interface type %T value %v\n",w,w)
}

func main(){
	p1 := person{
		name : "sandip",
	}
	var w worker =p1
	describe(w)
	w.work()
}
