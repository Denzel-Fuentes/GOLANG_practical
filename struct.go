package main

import "fmt"

/* struct (or structure) is a collection composed of fields, where
each field has a name and an associated type. Structures are used to
group related data to form logical records.
*/

//Definition
type Person struct {
	Name string
	Age  int
}

//methods
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

//Nested Structs
type Address struct {
    City, State string
}

type Admin struct {
    Name    string
    Age     int
    Address // Campo anónimo
}

func main() {

	p := Person{"Alice", 30}
	p2 := Person{Name: "Alice", Age: 30}
	p3 := new(Person)
	p3.Name = "Alice"
	p3.Age = 30

	//Anonymous structs
	var user struct {
		Username string
		Password string
	}
	user.Username = "user1"
	user.Password = "pass1"

	a := Admin{
		Name: "allice",
		Age: 30,
		Address: Address{
			City: "Santa c",
			State: "Ca",
		},
	}

	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(user)
	fmt.Println(p.Greet())
	fmt.Println(a)
	fmt.Println(a.City)

}
