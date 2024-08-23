package main
import(
	"fmt"
)
func main()  {
/* 	Go simplifies the use of loops compared to 
	other languages, providing only a basic loop structure: for
 */
	// Loop "For" (initialization ; condition ; post)
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i:", i)
	}
	
	//Simplified loop 
	i := 100
	for i < 105 {
		fmt.Println("Value of i=", i)
		i++
	}
	
	// Infinite loop
	/*for {
		fmt.Println("this is a infinite loop")
	} */
	
	//loop "for" with "range"
	
		//Iterating in a slice
	numbers := []int8{1, 2, 3, 4, 5}
	for index, value := range numbers {
    	fmt.Printf("index: %d, Value: %d\n", index, value)
	}
	
		//Iterating in a string
	word := "Golang"
	for index, char := range word {
    	fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

		//Iterating in a map
	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
		
	// "Break" and "Continue"
	for i:=0 ;i <10 ; i++{
		if i == 5{break} // Exit the loop when i is 5
		fmt.Println("i:", i)
	}

	for i:=0 ;i <10 ; i++{
		if i%2==0{continue} // Skip the rest of the loop if i is even
		fmt.Println("i:", i)
	}

	// Tags and Nested Loops 
	//Go allows you to label loops and refer to those labels to exit nested loops.
	outerLoop:
	for i := 0; i < 3; i++ {
    	for j := 0; j < 3; j++ {
        	if i == 1 && j == 1 {
            	break outerLoop  // // Exit both loops
        	}
        	fmt.Printf("i = %d, j = %d\n", i, j)
    	}
	}

	// Loops with  a one value in "range"
	/*If you only need the index or the value (but not both) 
	in a range loop, you can use _ to ignore the one you don't need.*/
	numbers2 := []int{1, 2, 3, 4, 5}
	for _, value := range numbers2 {
    	fmt.Println("Valor:", value)
	}
}