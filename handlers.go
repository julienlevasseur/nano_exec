package main

import (
	"net/http"
	"encoding/json"
)

// Define the index function
func Index(w http.ResponseWriter, r *http.Request) {
	//index_routes := []string{"providers"}
	var index_routes []string
	index_routes = GetJobs()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(index_routes); err != nil {
		panic(err)
	}
}
