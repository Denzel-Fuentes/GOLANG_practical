/* 
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE 754 32-bit floating-point numbers
float64     the set of all IEEE 754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32 
*/
package main
import(
	"fmt"
)
func main()  {
	//creating variables with var
	var name string = "Denzel"
	var number uint8 = 12
	var isChecked bool =  true

	//creation of variables without specifying the type
	var lastName = "Fuentes"
	var secondNumber = 15
	var isSelect = true

	//creation of variables without specifying the type and directly using ":="
	secondLastname := "Nunez"
	yearsOld := 20
	isAllowed := true
	
	
	fmt.Println("Hello world, I am", name , lastName, secondLastname)

	fmt.Println("My favorite number is",number, "and" , secondNumber , yearsOld)

	fmt.Println("isSelect:",isSelect,"\n"+
				"isAllowed:",isAllowed,"\n"+
				"isChecked:",isChecked )

	//Defaul values in variables
	var i int
    var f float64
    var b bool
    var s string
    var p *int
    var sl []int
    var m map[string]int
    var c chan int
    
	//"%v" print the value of a variable in its default representation, handles different types automatically.
	
    fmt.Printf("int: %v\n", i)       // 0
    fmt.Printf("float64: %v\n", f)   // 0
    fmt.Printf("bool: %v\n", b)      // false
    fmt.Printf("string: %v\n", s)    // ""
    fmt.Printf("pointer: %v\n", p)   // <nil>
    fmt.Printf("slice: %v\n", sl)    // []
    fmt.Printf("map: %v\n", m)       // map[]
    fmt.Printf("chan: %v\n", c)      // <nil>

	
}