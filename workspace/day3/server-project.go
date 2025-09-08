package main
import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
	"log"
)

type Orders struct {
	ID int `json:"id"`
	Item string `json:"item"`
	Quantity int `json:"quantity"`
}

func main() {
	var msg []byte
	var err error
	
	//orders in memory
	orders := []Orders {
		{ID: 1, Item: "Sandwich", Quantity: 3},
		{ID: 2, Item: "Coffee", Quantity: 32},
		{ID: 3, Item: "Muffin", Quantity: 13},
	} 

	// 2. HandleFunc
	http.HandleFunc("/orders", func(w http.ResponseWriter, req *http.Request) {
		
		if req.Method == "GET" {
			msg, err = json.Marshal(orders)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
		} else if req.Method == "POST" || req.Method == "DELETE" {
			body, err := io.ReadAll(req.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			var newOrder Orders
			err = json.Unmarshal(body, &newOrder)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if req.Method == "POST" {
				orders = append(orders, newOrder)
				log.Println(orders)
				msg = []byte("Appended: " + string(body))
			} else if req.Method == "POST" {
				for i, o := range orders {
					if o.ID == newOrder.ID {
						orders = append(orders[:i], orders[i+1:]...)
						msg = []byte("Deleted: " + string(body))
						break
					}
				}
			}
			msg = []byte("ID not found to be Deleted: " + string(body))
			log.Println(orders)
		} else {
			fmt.Println("UNKNOWN METHOD")
		}
		
		io.WriteString(w, string(msg))
	})
	
	// 1. Start server on localhost:8080. ListenAndServe
	log.Println("Server running. Listening on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if (err != nil) {
		log.Fatal(err)
	}
}

// How to connect to server:
// GET: curl http://localhost:8080/orders
// POST: curl -X POST http://localhost:8080/orders -d '{"id":1,"item":"Burger","quantity":2}' -H "Content-Type: application/json"
// DELETE: curl -X DELETE http://localhost:8080/orders -d '{"id":41}'
