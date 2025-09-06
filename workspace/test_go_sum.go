package main

import "github.com/google/uuid"

func main() {
	id := uuid.New()
	fmt.Println("Generated UUID:", id)
}
