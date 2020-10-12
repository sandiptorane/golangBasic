package main

import "fmt"

type author struct {
	firstname string
	lastname string
	bio string
}

type post struct{
	title string
	content string
	author         //author is embemdded into post
}

func (a author)fullName() string{
	return fmt.Sprintf("%s %s",a.firstname,a.lastname)
}

func (p post)details(){
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

func main(){
	author1 := author{
		"Sandip",
		"torane",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()
}
