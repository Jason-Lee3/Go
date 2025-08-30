package main

import (
	"fmt"
	"net/http"
)

func main() {
	// tasks := []task{}

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
	fmt.Println("Received a request!")
}
