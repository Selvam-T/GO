package main

import (
	"fmt"
	/*"encoding/json"*/
)

func main() {
	/* if else */
	something := map[string]int{"Number": 200}
	if (something["Number"] == 1200){
		fmt.Printf("Number %v\n",something["Number"])
	} else {
		fmt.Printf("Else it is not\n")
	}
	
	/* Type switch */
	var num interface{} = 4
	switch num.(type) {
		case int:
			fmt.Println("int baby")
		case string:
			fmt.Println("string baby")
		default:
			fmt.Println("Unknown type baby")
	}
	
	/* Basic switch */
	x := 300

	switch {
		case (x % 2 == 0):
			fmt.Println("Even number")
		default:
			fmt.Println("Odd number")
	}

	/* Multiple values in switch  */
	switch x {
		case 1, 2, 3:
			fmt.Println("case 1, 2, 3")
		case 10, 20, 30:
			fmt.Println("case 10, 20, 30")
		case 100, 200, 300:
			fmt.Println("case 100, 200, 300")
	}

}
