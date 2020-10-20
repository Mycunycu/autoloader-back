package models

import "net/http"

// Route - struct for each new route
type Route struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}
