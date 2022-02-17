package main

// Imports log and net/http packages from std lib
import (
	"log"
	"net/http"
)

// Defines a home handler function who writes a byte slice as response body
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi from Snippetbox"))
}

// App entry point
func main() {
	// Initializing servermux and  register homeHandler function as route "/" handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	/*
		Initializing http server with http.ListAndServe function. Thats receive two parameters:
			- TCP address to list (in host:port format)
			- application servemux
		ListAndServe function return a err value
	*/
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	// Logs error and exit
	log.Fatal(err)
}
