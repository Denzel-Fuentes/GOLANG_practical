package main

import "fmt"

/* An interface is a type that specifies a set of methods, but does not
implement these methods itself. Interfaces are used to define general
behaviors that can be shared and adopted by multiple concrete types. They provide a level of abstraction and allow polymorphic programming in Go
*/

//Definition
type Speaker interface {
	Speak() string
}

//Implementation
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

//use
func MakeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

//void interface
func PrintAnything(v interface{}) {
    fmt.Println(v)
}


func main() {
	d := Dog{
		Name: "dani",
	}
	c := Cat{Name: "Whiskers"}
	MakeItSpeak(c)
	fmt.Println(d.Speak())
}
