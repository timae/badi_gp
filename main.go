package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Flussbad struct {
    Name         string `xml:"title"`
    Capacity     string `xml:"capacity"`
    Status       string `xml:"status"`
    LastModified string `xml:"lastModified"`
    URL          string `xml:"url"`
}

type FlussbadData struct {
    XMLName xml.Name   `xml:"flussbadData"`
    Flussbads []Flussbad `xml:"flussbad"`
}

func main() {
    // Define a handler function
    handler := func(w http.ResponseWriter, r *http.Request) {
        // Fetch data from XML API
        resp, err := http.Get("https://data.stadt-zuerich.ch/dataset/wassertemperaturen-freibaeder/resource/548d1ceb-1daf-4cf9-a14a-92c86326824d")
        if err != nil {
            http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        // Read response body
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            http.Error(w, "Failed to read response body", http.StatusInternalServerError)
            return
        }

        // Parse XML
       // Parse XML
        var data FlussbadData
        if err := xml.Unmarshal(body, &data); err != nil {
            http.Error(w, fmt.Sprintf("Failed to parse XML: %v", err), http.StatusInternalServerError)
            return
        }


        // Display data in the response
        for _, flussbad := range data.Flussbads {
            fmt.Fprintf(w, "Name: %s\n", flussbad.Name)
            fmt.Fprintf(w, "Capacity: %s\n", flussbad.Capacity)
            fmt.Fprintf(w, "Status: %s\n", flussbad.Status)
            fmt.Fprintf(w, "Last Modified: %s\n", flussbad.LastModified)
            fmt.Fprintf(w, "URL: %s\n", flussbad.URL)
            fmt.Fprintln(w, "------")
        }
    }

    // Register the handler function to handle all requests to the root URL ("/")
    http.HandleFunc("/", handler)

    // Start the HTTP server on port 8080
    fmt.Println("Server listening on :8080")
    http.ListenAndServe(":8080", nil)
}
