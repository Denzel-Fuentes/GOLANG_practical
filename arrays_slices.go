package main

import "fmt"

func main()  {

	//Ways to create arrays
	var arrayNumbers [5]int
	var listFruits = [4] string {"Pera","Manzana","Naranja"}
	listPaises := [4]string{"Peru","Bolivia","Ecuador","Brasil"}
	listNames :=[]string{"Denzel","Daniel","Name2"}

	//print list of fruits in position 2
	fmt.Println(listFruits[2])
	fmt.Println(arrayNumbers[3])

	//change the value of a position
	listPaises[0] = "Chile"
	fmt.Println(listPaises)

	//add a new value to the array
	listNames = append(listNames,"Gabriel")
	fmt.Println(listNames)

	//cut a part of the array in this case from position 1 to position 2, it will not take the value that is in position 3 only until position 2
	listPaises2 := listPaises[1:3]
	fmt.Println(listPaises2) // [Bolivia Ecuador]


	//takes all values ​​below position 2
	listPaises3 := listPaises[:2]
	fmt.Println(listPaises3) //[Chile Bolivia]

	//takes all values ​​greater than and equal to position 1
	listPaises4 := listPaises[1:]
	fmt.Println(listPaises4)//[Bolivia Ecuador Brasil]
}
