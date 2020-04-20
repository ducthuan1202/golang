package main

import (
	"fmt"
)

type Person interface {
	Run()
	Eat()
}

type Human struct {
	Name string
	Age int
}

func (h Human) String()string{
	return fmt.Sprintf("{%s: %d}", h.Name, h.Age)
}

func (h Human) Run() {
	fmt.Println("Human is run")
}

func (h Human) Eat() {
	fmt.Println("Human is eat")
}

func (h Human) GetName() string {
	return h.Name
}

func main() {

	var p Person

	p = Human{
		Name: "Nguyen Duc THuan",
		Age: 30,
	}

	fmt.Println(p)

	h := Human {
		Name: "Trang",
		Age: 28,
	}
	p = h
	fmt.Println(h.GetName())
}
