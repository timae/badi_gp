package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!") // Write response to the client
    }

    // Register the handler function to handle all requests to the root URL ("/")
    http.HandleFunc("/", handler)

    // Start the HTTP server on port 8080
    fmt.Println("Server listening on :8080")
    http.ListenAndServe(":8080", nil)
}
