package main

import "fmt"

type animal interface {
	print()
}

type dog struct{}

type puppy struct {
	*dog
}

func (d *dog) print() {
	fmt.Println("Dog")
}

func main() {

	// p := puppy{&dog{}}
	// var a animal = p // this will work

	// var a animal = puppy{} // this will work
	var a animal = &puppy{} // this also will work
	a.print()

	// var b animal = dog{} // this will not work
	var b animal = &dog{}
	b.print()
}

/*
Output: Dog

Explanation:
When we assign a puppy struct to an animal interface, the interface value is a pointer to a nil interface value.
When we call the print method on the interface, it will print "Dog" because the print method is a pointer receiver.

*/
