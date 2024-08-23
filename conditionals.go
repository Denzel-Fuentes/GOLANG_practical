package main
import(
	"fmt"
)
func main()  {
	//Conditional "IF - ELSE "
	var yearsOld int8 = 39
	if yearsOld >= 60 {
		fmt.Println("this person can retire")
	}else if yearsOld >= 18 {
		fmt.Println("this person cannot retire")
	}else{
		fmt.Println("the person is a minor")
	}

	//Initialization inside if
	var x int8 = 10
	if y := x + 1; y > 10 {
		fmt.Println("y is greater than 10")
	}
	
	//Conditional "Switch"
	var dayOfWeek string = "Friday"
	switch dayOfWeek {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Almost the weekend")
	default:
		fmt.Println("Another day of the week")
	}

	//Multiple conditions in a "case"
	dayOfWeek = "Saturday"
	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("End of the week")
	default:
		fmt.Println("It's a working day")
	}
	
	//Switch without variable , to evaluate multiple conditions.
	x = -120
	switch {
	case x > 100:
		fmt.Println("x is greater than 100")
	case x < 0:
		fmt.Println("x is a negative number")
	default:
		fmt.Println("x is between 0 and 100")
	}
	
	//Use of "fallthrough", so that control passes to the next case
	switch x {
	case 1:
		fmt.Println("x is 1")
		fallthrough
	case 2:
		fmt.Println("x is 2 or 1")
	default:
		fmt.Println("x is neither 1 nor 2")
	}
	
	fmt.Println("end program") 
}