package main

/* A pointer is a variable that stores the memory address of another variable.
Pointers are very useful in Go for various tasks, including modifying data
in functions or efficiently manipulating large data structures without copying them. */


import (
	"fmt"
)

// & = para conocer la direccion de memoria donde esta almacenado la variable
// * = para acceder al valor que se encuentra almacenado en la direccion de memoria

//use of pointer in functions
func increment(x *int) {
    *x = *x + 1
}

type Person4 struct {
    Name string
    Age  int
}

func main()  {
	//declaration
	//var ptr *int 	  //pointer to a int
	var a int = 58
	var ptr *int = &a //"&"" is used to get the memory address of a variable

	//Pointer dereferencing
	var value int = *ptr //This assigns the value of the variable pointed to by ptr (in this case a) to the variable value.

	b := 10
	increment(&b)
	fmt.Println(b)
	fmt.Println(value)

	p := &Person4{"John", 30} //we get the memory address
    fmt.Println(p.Name)  // Direct access to the structure fields through the pointer

}