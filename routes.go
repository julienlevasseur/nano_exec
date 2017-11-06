package main

import (
	"net/http"
)

// Route definition
type Route struct {
	Name		string
	Method		string
	Pattern		string `json:"pattern"`
	HandlerFunc http.HandlerFunc
}

// Routes
var routes []Route = []Route{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},
}
