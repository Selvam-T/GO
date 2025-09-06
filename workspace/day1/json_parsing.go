package main

import (
	"encoding/json"
	"fmt"
)

// User represents our internal Go data model.
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Secret    string `json:"-"` // This field will be ignored
	IsActive  bool   `json:"is_active,omitempty"`
}

func main() {
	// 1. Marshal (Serialize): Go struct to JSON byte slice
	fmt.Println("--- Marshaling ---")

	// Create a User struct with a last name and an active status
	user1 := User{
		FirstName: "Jane",
		LastName:  "Doe",
		Secret:    "secret-password", // This field will be ignored
		IsActive:  true,
	}

	// The custom JSON keys and ignored field will be used here.
	jsonData1, _ := json.Marshal(user1)
	fmt.Println("Original User:", string(jsonData1))

	// Create a User struct without a last name.
	// `omitempty` will cause the Last Name and Is Active fields to be omitted.
	user2 := User{
		FirstName: "John",
		Secret:    "another-secret",
	}
	jsonData2, _ := json.Marshal(user2)
	fmt.Println("User with empty fields:", string(jsonData2))

	// 2. Unmarshal (Deserialize): JSON byte slice to Go struct
	fmt.Println("\n--- Unmarshaling ---")

	jsonString := `{"first_name":"Alice","last_name":"Smith","is_active":true}`
	var alice User

	// Unmarshal the JSON byte slice into our struct.
	// Use []byte(jsonString) to convert the string to a byte slice.
	err := json.Unmarshal([]byte(jsonString), &alice)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}
	// The Go struct field names will be used for access.
	fmt.Printf("Go struct fields from JSON: FirstName=%s, LastName=%s, IsActive=%t\n", alice.FirstName, alice.LastName, alice.IsActive)

	// Since "Secret" was ignored during marshaling, it will be its zero value (empty string) after unmarshaling.
	fmt.Printf("Secret field is not unmarshaled: '%s'\n", alice.Secret)
}
