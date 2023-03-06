package main

import (
	"log"
	"net/http"
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
	writer.Write([]byte("Display a specific snippet..."));
}

func createSnippet (writer http.ResponseWriter, request *http.Request) {
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