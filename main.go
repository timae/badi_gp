package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define struct to represent the XML data structure
type BathInfo struct {
	XMLName xml.Name `xml:"bathinfos"`
	Baths   []Bath   `xml:"baths>bath"`
}

type Bath struct {
	Title  string `xml:"title"`
	Poiid  string `xml:"poiid"`
	// Add more fields if needed
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Read XML data from the URL
		xmlData, err := readXML("https://www.stadt-zuerich.ch/stzh/bathdatadownload")
		if err != nil {
			http.Error(w, "Failed to fetch XML data", http.StatusInternalServerError)
			return
		}

		// Parse XML data
		bathInfo := BathInfo{}
		err = xml.Unmarshal(xmlData, &bathInfo)
		if err != nil {
			http.Error(w, "Failed to parse XML data", http.StatusInternalServerError)
			return
		}

		// Generate HTML response with dropdown menu
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		    <title>Bath Information</title>
		    <script>
		        function onSelectChange() {
		            var selectedPoiid = document.getElementById("bathDropdown").value;
		            var selectedTitle = document.getElementById("bathDropdown").options[document.getElementById("bathDropdown").selectedIndex].text;
		            document.getElementById("selectedPoiid").textContent = "Selected POIID: " + selectedPoiid;
		            document.getElementById("selectedTitle").textContent = "Selected Title: " + selectedTitle;
		        }
		    </script>
		</head>
		<body>
		    <h2>Select a Bath</h2>
		    <select id="bathDropdown" onchange="onSelectChange()">
		`

		// Populate dropdown menu options
		for _, bath := range bathInfo.Baths {
			html += fmt.Sprintf(`<option value="%s">%s</option>`, bath.Poiid, bath.Title)
		}

		// Close HTML tags and display selected values
		html += `
		    </select>
		    <div id="selectedPoiid"></div>
		    <div id="selectedTitle"></div>
		</body>
		</html>
		`

		// Write HTML response
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	http.ListenAndServe(":8080", nil)
}

// Function to read XML data from URL
func readXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
