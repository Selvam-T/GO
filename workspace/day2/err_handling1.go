package main
import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(100,10)
	if err != nil {
		fmt.Println("Error: ",err)
	}
	fmt.Println("Result: ",result)
}
