package main

import "fmt"

type User struct {
	Name string
}

func (u User) changeName(name string) {
	u.Name = name
}

func main() {

	u := User{Name: "Nick"}
	var x interface{} = u

	user := x.(User)
	user.changeName("John")

	fmt.Println(u.Name)
	fmt.Println(user.Name)
}

/*
Output:
Nick
Nick

Why?
This is a combination of interface+struct+value reciever.

When we assign a User struct to an interface, the interface value is a pointer to a nil interface value.
When we try to call the changeName method on the interface, the method is called on the copy of the User struct.
So the original User struct is not changed.

If we want to change the original User struct, we need to use a pointer receiver.

*/
