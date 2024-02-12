package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home (writer http.ResponseWriter, request *http.Request) {
	// Since we're serving the home handler to ServeMux in a sub tree, we
	// can instead do our check here to make sure we're serving the correct URL
	if request.URL.Path != "/" {
		http.NotFound(writer, request);
		return;
	}

	// We cast the string to a byte array.
	writer.Write([]byte("Hello from Snippetbox"));
}

func showSnippet (writer http.ResponseWriter, request *http.Request) {
	// Extract the value of the id parameter from the URL string
	// Try to convert it to an integer using the strconv.Atoi package
	// If it cant convert it to an integer or the value is less than 1
	// Return a 404
	id, err := strconv.Atoi(request.URL.Query().Get("id"));
	if (err != nil || id < 1) {
		http.NotFound(writer, request);
		return;
	}
	writer.Write([]byte("Display a specific snippet..."));

	// Use the fmt.Fprintf() function to interpolate the ID value
	// and include it in our response
	fmt.Fprintf(writer, "Displaying a specific snippet with ID %d...", id);
}

// First implementation of createSnippet using raw writer.WriteHeader
// func createSnippet (writer http.ResponseWriter, request *http.Request) {
// 	// Restrict our createSnippet method to just POST requests
// 	if request.Method != "POST" {
// 		// It's only possible to call WriteHeader once per handler.
// 		// After the status code has been written, it cannot be changed.
// 		writer.WriteHeader(405);
// 		writer.Write([]byte("Method not allowed"));
// 		return;
// 	}

// 	writer.Write([]byte("Create a new snippet..."));
// }

// Second implementation that uses the helper http.Error
// This calls writer.WriteHeader under the hood
func createSnippet (writer http.ResponseWriter, request *http.Request) {
		// Restrict our createSnippet method to just POST requests
		if request.Method != "POST" {
			// Use the http.Error method to send a 405 status code
			// with the error message of Method not allowed
			http.Error(writer, "Method not allowed", 405);
			return;
		}
	
		writer.Write([]byte("Create a new snippet..."));
	}

func main () {
	// Initialize a new ServeMux and set it to our var mux
	// A ServeMux is a router in Go
	// The type is inferred because we use the shorthand :=
	mux := http.NewServeMux();

	// Assign the handler function to our home page within our router (ServeMux)
	// This is a subtree path example because it ends in a trailing slash.
	// You can think of it ending in a wildcard like '/**'
	mux.HandleFunc("/", home);
	// These are two examples of fixed paths -- paths that don't end with a trailing slash.
	// These handlers will only be called when the URL path matches exactly.
	mux.HandleFunc("/snippet", showSnippet);
	mux.HandleFunc("/snippet/create", createSnippet);


	log.Println("Starting the server on :4000");
	// Start up our web server on port 4000
	// Pass in our ServeMux as our router for the webserver
	// If ListenAndServe returns an error we'll log it and exit the app
	err := http.ListenAndServe(":4000", mux);

	// log.Fatal will also call os.Exit(1) after it prints its log message
	// immediately exiting the application.
	log.Fatal(err);
}