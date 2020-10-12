package main

import "fmt"

type person struct{
	name string
	age int
}

type Address struct {
	state string
	country string
}

type describer interface {
	describe()
}

func (p person)describe(){   //implemented using value receiver it also  revceives pointer
	fmt.Printf("name = %s age =%d\n",p.name,p.age)
}

func (a *Address)describe(){  //implemented using pointer receiver  it only receives pointer
	fmt.Printf("Address  state = %s coutry = %s",a.state,a.country)
}

func main(){
	var d1 describer
	p1 := person{
		name: "sandip torane",
		age : 23,
	}
	d1 = p1 //call by value
	d1.describe()
	p2 :=person{"dinesh melge",23}
	var d2 describer
	d2= &p2    //d2 points to p2
	d2.describe()


	a := Address{"maharashtra","India"}
	d1 = &a       //d1 points to a
	d1.describe()

}
