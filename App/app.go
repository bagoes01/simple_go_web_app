package main
 
import (
"fmt"
"net/http"
)
 
func main() {
	// Register the handler function for the root route
    http.HandleFunc("/", indexPage)
		// Start the server and listen on port 8000
    	http.ListenAndServe(":8080", nil)
}

// indexPage is the handler function for the root route "/"
func indexPage(w http.ResponseWriter, r *http.Request) {

    //do anything

	// Write the response
	fmt.Fprintf(
        w,
		`
		<h1>Hello World!</h1>
    	<p>Hello World!</p>
		`)
	fmt.Fprintf(w, "Welcome to a page about ", r.URL.Path)
	fmt.Fprintf(w, "Welcome to a page about ", r.URL.Path[1:])
}
