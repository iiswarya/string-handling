package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Dog")
}

func main() {

	var a Animal = Dog{}
	d, ok := a.(*Dog)

	fmt.Println(d, ok)
}

/*
Output: <nil>,false

Explanation:
When we assign a Dog struct to an Animal interface, the interface value is a pointer to a nil interface value.
When we try to assert the interface value to a Dog struct, the type assertion fails because the interface value is a pointer to a nil interface value.
Important thing to note here we need to use comma-ok idiom to check if the type assertion is successful. So if assertion is failed we will recieve a nil value and false
If we dont use comma-ok idiom, then we will get a panic.



*/
