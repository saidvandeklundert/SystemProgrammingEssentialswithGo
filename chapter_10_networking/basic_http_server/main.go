package main

/*
	http://127.0.0.1:8080/
*/

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Handling a GET request on path %s\n", r.URL.Path)
	case http.MethodPost:
		fmt.Fprintf(w, "Handling a POST request on path %s\n", r.URL.Path)
	case http.MethodPut:
		fmt.Fprintf(w, "Handling a PUT request on path %s\n", r.URL.Path)
	case http.MethodDelete:
		fmt.Fprintf(w, "Handling a DELETE request on path %s\n", r.URL.Path)
	default:
		http.Error(w, "Unsupported HTTP method", http.
			StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
