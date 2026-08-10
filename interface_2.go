package main

import "fmt"

type Dog struct{}

type Animal interface {
	Bark()
}

func (d *Dog) Bark() {
	fmt.Println("Woof!")
}

func main() {

	var d *Dog = nil

	var a Animal = d

	fmt.Println(a == nil)

	a.Bark()

}

// Explanation
/*
In go, interface will take a key and value pair
So here key is dog and value is nil
if key and value both nil, then the interface is nil
Either key or value is not nil, then the interface is not nil
so a==nil is false and it would print the bark method
*/
