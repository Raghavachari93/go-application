package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go running inside Docker on AWS EC2!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
