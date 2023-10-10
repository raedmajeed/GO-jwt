package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error listening to port 8080")
		return
	}
}