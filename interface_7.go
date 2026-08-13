package main

import "fmt"

type User struct {
	Name string
}

type Printer1 interface {
	print()
}

func (u *User) print() {
	fmt.Println(u.Name)
}

func main() {

	var u *User = nil
	var p Printer1 = u

	fmt.Println(p == nil)
	p.print()
}

/*
Output: false and p.print() will panic

Explanation:
In go, when we assign a nil value to an interface, the interface is not nil.
So p == nil is false.
When we call the print method on the interface, it will panic because the interface is nil and the print method is a pointer receiver.

*/
