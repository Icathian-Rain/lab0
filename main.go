package main

import (
	"fmt"
	"main/handlers"
	"net/http"
)

func main() {
	// TODO: some code goes here
	// Fill out the HomeHandler function in handlers/handlers.go which handles the user's GET request.
	// Start an http server using http.ListenAndServe that handles requests using HomeHandler.
	fmt.Println("start server")
	http.HandleFunc("/", handlers.HomeHandler)
	http.ListenAndServe("127.0.0.1:6670", nil)
}
