package main

import ( "fmt")

func main() {
	j := 0
	/* for loop */
	for i := 0; i < 2; i++ {
		/* while loop */
		for j < 3 {
			fmt.Println(i == j)
			j++
		}
	}
	/* Infinite loop */
	for {
		if (j == 5) {
			fmt.Println("j == 5")
			break;
		}
		j++
	}
	
}
