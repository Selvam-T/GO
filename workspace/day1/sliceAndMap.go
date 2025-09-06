package main

import "fmt"

func main() {
	// --- Slices ---
	fmt.Println("--- Slices (Dynamic Arrays) ---")
	// Creating a slice with a literal
	numbers := []int{10, 20}
	fmt.Println("Initial slice of numbers:", numbers)

	// Appending to a slice
	numbers = append(numbers, 30, 40)
	fmt.Println("Slice after appending:", numbers)

	// Iterating over a slice
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("Iterating over slice of names:")
	for i, name := range names {
		fmt.Printf("  Index %d: %s\n", i, name)
	}

	// --- Maps ---
	fmt.Println("\n--- Maps (Key-Value Pairs) ---")
	// Map with string keys and integer values
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("Initial map with string keys:", ages)

	// Accessing and checking for a key
	if age, ok := ages["Alice"]; ok {
		fmt.Println("  Alice's age (found):", age)
	}
	if age, ok := ages["Diana"]; ok {
		fmt.Println("  Diana's age (found):", age)
	} else {
		fmt.Println("  Diana not found in map.")
	}

	// Delete a key
	delete(ages, "Bob")
	fmt.Println("Map after deleting 'Bob':", ages)

	// Map with integer keys and string values (addressing your question)
	statusCodeMessages := map[int]string{
		200: "OK",
		404: "Not Found",
		500: "Internal Server Error",
	}
	fmt.Println("\nMap with integer keys:", statusCodeMessages)
	fmt.Println("Status for 200:", statusCodeMessages[200])

	// Iterating over a map (order is not guaranteed)
	fmt.Println("Iterating over map with string keys:")
	for name, age := range ages {
		fmt.Printf("  Name: %s, Age: %d\n", name, age)
	}
}
