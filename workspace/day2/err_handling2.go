package main
import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json:"first_name"`
	Age int `json:"age"`
	Gender string `json:"gender,omitempt"`
	Email string `json:"-"`
}

func main() {
	//struct instance
	data := User{
		FirstName: "Jane",
		Age: 20,
		Gender: "M",
		Email: "dummy",
	}

	jsonData, err := json.Marshal(data)
	if (err == nil){
		fmt.Println(string(jsonData))
	}

	//define a map
	data2 := map[string]interface{} {
		"name": "Aura",
		"age": 200,
		"phone": "12345",
		"func": func() {},	//Fucntion can't be marshaled
	}
	jsonData2, err := json.Marshal(data2)
	if (err != nil) {
		fmt.Println("Error Marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData2))
}
