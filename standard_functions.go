package main

import (
	"fmt"
)

// Declaration Functions
func sumar(a int, b int) int {
	 return a + b
}

//Parameters
func sumaVarios(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

//Return of funtion
func dividir(a,b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("no se puede dividir por cero")
    }
    return a / b, nil
}

//anonymous functions
var cuadra = func(x int )int {
    return x*x
}

//Closures
/* Closures in programming refer to a function that "remembers" 
the environment or context in which it was created, even after 
said context has disappeared.
 */
func sequencer() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

//recursion
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

//function map
var functions = map[string] func(int,int) int{
    "sum":       func(i1, i2 int) int {return i1+i2},
    "subtract" : func(i1, i2 int) int {return i1-i2},
}

func main() {
	resultadoSuma := sumar(3, 5)
	fmt.Println("Resultado de sumar 3 y 5:", resultadoSuma)

	resultadoSumaVarios := sumaVarios(1, 2, 3, 4)
	fmt.Println("Resultado de sumar varios números (1, 2, 3, 4):", resultadoSumaVarios)

	resultadoDividir, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error al dividir:", err)
	} else {
		fmt.Println("Resultado de dividir 10 por 2:", resultadoDividir)
	}
	_, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Error al dividir por cero:", err)
	}

	resultadoCuadrado := cuadra(4)
	fmt.Println("El cuadrado de 4 es:", resultadoCuadrado)

	next := sequencer()
	fmt.Println("Primera llamada a sequencer:", next())
	fmt.Println("Segunda llamada a sequencer:", next())
	fmt.Println("Tercera llamada a sequencer:", next())

	resultadoFactorial := factorial(5)
	fmt.Println("El factorial de 5 es:", resultadoFactorial)

    resultadoMap := functions["sum"](1,2)
    fmt.Println(resultadoMap) 
}