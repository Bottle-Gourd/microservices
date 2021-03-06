// Marshalling and un-marshaling JSON
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type helloWorldRequest struct {
	Name string `json:"name"`
}
type helloWorldResponse struct {
	Message string `json:"message"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	response := helloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// NOTE: run the following curl command to test the above code
// curl localhost:8080/helloworld -d '{"name": "There"}'
