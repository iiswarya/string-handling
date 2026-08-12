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
