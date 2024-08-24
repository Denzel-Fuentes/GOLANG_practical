package main
import(
	"fmt"
)

// Declaration of Functions
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
func dividir(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("no se puede dividir por cero")
    }
    return a / b, nil
}

func main()  {
	suma1 := sumaVarios(1, 2)
    fmt.Println("La suma de 1 y 2 es:", suma1)
	fmt.Printf("End of program")
}