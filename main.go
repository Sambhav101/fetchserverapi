package main

import (
	"encoding/json"
	points "fetchapi/logic"
	"fetchapi/receipt"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// global variable to store data in memory
var receiptPoints = make(map[string]int)

func processReceipts(w http.ResponseWriter, r *http.Request) {
	// decode the request body and process the receipt
	var receiptMetaData receipt.Receipt
	err := json.NewDecoder(r.Body).Decode(&receiptMetaData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// calculate total points
	totalPoints := points.CalculatePoints(receiptMetaData)

	// generate a unique id for each receipt
	id := uuid.New().String()

	// store the result in our variable
	receiptPoints[id] = totalPoints

	// create a response for the post request
	response := receipt.ReceiptId{ID: id,}

	// send the response back to the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	// obtain the variable id from the route
	vars := mux.Vars(r)
    id := vars["id"]

	// check if the id exists in memory
	points, exists := receiptPoints[id]
	if !exists {
		http.NotFound(w, r)
		return
	}
		
	// create a response for the get request
	response := receipt.TotalPoints{TotalPoints: points,}

	// send the response back to the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response) 
}

func main() {

	// Create a new Gorilla Mux router
    r := mux.NewRouter()

	// defining routes
	r.HandleFunc("/receipts/process", processReceipts).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", getPoints).Methods("GET")

	// listen on port 8080
	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8080", r)
}



