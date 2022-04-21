package gohelloworld

import (
	"net/http"
	"fmt"
)

// Web server (Listen on port 8080) and call
func Webserver() {
	http.HandleFunc("/hello", Handler)
	http.ListenAndServe(":8080", nil)
}

// Handler function
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Hello())
}
