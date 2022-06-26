// Marshalling and un-marshaling JSON
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type helloWorldResponse struct {
	//Message change the output field to be 'message'
	Message string `json:"message"`
	//Author do not output this field
	Author string `json:"-"`
	// Date do not output if the field value is empty
	Date string `json:",omitempty"`
	// Id rename to id
	Id int `json:"id"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{
		Message: "HelloWorld!",
		Author:  "Dave",
		Date:    time.Now().String(),
		Id:      777,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Print("Error occurred while encoding: ", err)
	}
}
