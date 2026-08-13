//interface + embedding + shadowing

package main

import "fmt"

type speaker interface {
	speak()
}

type animal struct{}

func (a *animal) speak() {
	fmt.Println("Animal")
}

type dog struct {
	animal
}

func (d *dog) speak() {
	fmt.Println("Dog")
}

func main() {

	d := dog{}
	d.speak()
	d.animal.speak()
}

/*

Output: Dog
Animal

Explanation:
When we call the speak method on the dog struct, it will print "Dog"
When we call the speak method on the animal struct, it will print "Animal" because its shadowed by the dog struct.

*/
