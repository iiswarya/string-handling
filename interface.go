package main

import "fmt"

type Printer interface {
	Print()
}

type User struct {
	Name string
}

func (u *User) Print() {
	fmt.Println("Printing username:", u.Name)
}

func main() {

	//var p Printer = User{Name: "John"} - will this work?
	//no, it will not work because the User struct does not implement the Printer interface.
	//to fix this, we need to change the Print method to a pointer receiver.

	var p Printer = &User{Name: "John"}
	p.Print() // this will work because the Print method is a pointer receiver.
}
