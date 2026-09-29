package main

import (
	"os"
	"fmt"
)

type Person struct{
	firstName string
	lastName string
}

func main() {
	p := new(Person)
	p.firstName = "Kevin Hayes"
	p.lastName = "Anderson"
	person := *p
	fmt.Println(person)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	ns := names[1:]
	for i:=range ns {
		fmt.Println(ns[i])
	}
	os.Exit(0)
}