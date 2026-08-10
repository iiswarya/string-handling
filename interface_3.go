package main

import "fmt"

type Engine struct{}

type Car struct {
	Engine // Composition (Embedded struct)
	Model  string
}

func (e *Engine) start() {
	fmt.Println("Engine started")
}

func main() {

	c := Car{Model: "Tesla"}
	c.start()
}

/*
output: Engine started

Explanation:
In go, When we embed a struct, all Engine methods are available on the car struct - This is called method promotion

*/
