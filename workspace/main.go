package main

import (
	"encoding/json"
	"fmt"
)

type User struct {

	/*Struct Field name vs cutome JSON keyname*/
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email_address"`
}

func main() {
	/*unmarshal: parsing a JSON string, convert to Go data struct*/
	jsonData := []byte(`{"name": "Alice", "age": 30, "city": "New York"}`)
	
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for  key, value := range data {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}

	/*marshal: convert Go data struct into a JSON string */
	user := User{FirstName: "Alex", Email: "Alex@example.com"}
	/*_ is blank identified to discard 2nd return value, error */
	jsonData1, _ := json.Marshal(user)
	fmt.Println(string(jsonData1))
}
