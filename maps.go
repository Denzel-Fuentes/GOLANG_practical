package main
import "fmt"


/* In Go (Golang), maps are built-in data structures that associate keys 
with values, allowing you to store and retrieve data efficiently. */

func main() {
	//Declaration and Initialization of Maps (  map[key]value  )
    
	//1st option
	myMap := map[string]string{
		"Colombia":"Bogota",
		"Chile":"Santiago",
		"Venezuela " : "Caracas",
	}
	myMapNumberString := map[int]string{
		1:"Bogota",
		2:"Santiago",
		3: "Caracas",
	}
	myMapStringNumber := map[string]int16{
		"Bogota":1,
		"Santiago":342,
		 "Caracas":432,
	}

	//2st option 
		// declaration
	var myMap2 map[string]int
		//Initialization, "make" allocates memory to start working with the map
	myMap2= make(map[string]int) // or myMap := make(map[string]int)

	//Nested Maps
	nestedMap := make(map[string]map[string]int)
	nestedMap["fruits"] = map[string]int{
		"apple":  2,
		"banana": 5,
	}
	

	//Add and Update Elements
	myMap["Chile"] = "santiago de chile"
	myMap["Bolivia"] = "Sucre"

	//Delete Elements
	delete(myMap,"Venezuela")

	fmt.Println("The map is:",myMap)

	//Recover Values map[key]
	fmt.Println("La capital de chile es:",myMap["Chile"]) //"santiago de chile"
	fmt.Println(myMap2)
	fmt.Println(myMapNumberString)
	fmt.Println(myMapStringNumber)

}