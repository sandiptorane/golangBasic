package main

import "oop/new_employee"

func main(){
	e := new_employee.New("sandip","torane",30,5)
	e.CalculateLeaves()
}
