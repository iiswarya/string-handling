package main

import "fmt"

type Engine struct{}

func (e Engine) start() {
	fmt.Println("Engine started")
}

type ElectricCar struct {
	Engine
}

func (e *ElectricCar) start() {
	fmt.Println("Electric car started")
}

func main() {

	eCar := ElectricCar{}
	eCar.start()
}

/*
What will be the output?
	Electric car started

Explanation:
In go, when we embed a struct, all Engine methods are available on the ElectricCar struct - This is called method promotion
so when we call eCar.start(), it will call the start method of the ElectricCar struct, because ElectricCar defines its own start method. This shadows the promoted method
if we call explicitly eCar.Engine.start(), then it will point to the start methos of the Engine struct
*/
