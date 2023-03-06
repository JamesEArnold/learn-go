package main

import (
	"log"
	"net/http"
)

func home (writer http.ResponseWriter, request *http.Request) {
	// We cast the string to a byte array.
	writer.Write([]byte("Hello from Snippetbox"))
}

func main () {
	// Initialize a new ServeMux and set it to our var mux
	// A ServeMux is a router in Go
	// The type is inferred because we use the shorthand :=
	mux := http.NewServeMux();

	// Assign the handler function to our home page within our router (ServeMux)
	mux.HandleFunc("/", home);


	log.Println("Starting the server on :4000");
	// Start up our web server on port 4000
	// Pass in our ServeMux as our router for the webserver
	// If ListenAndServe returns an error we'll log it and exit the app
	err := http.ListenAndServe(":4000", mux);

	// log.Fatal will also call os.Exit(1) after it prints its log message
	// immediately exiting the application.
	log.Fatal(err);
}