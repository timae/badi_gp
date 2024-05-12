package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BathInfo struct {
	Title           string `xml:"title"`
	TemperatureWater string `xml:"temperatureWater"`
	OpenClosedText  string `xml:"openClosedTextPlain"`
	URLPage         string `xml:"urlPage"`
}

type BathInfos struct {
	XMLName xml.Name   `xml:"bathinfos"`
	Baths   []BathInfo `xml:"baths>bath"`
}

func main() {
	// Make a GET request to the URL
	resp, err := http.Get("https://www.stadt-zuerich.ch/stzh/bathdatadownload")
	if err != nil {
		fmt.Println("Failed to fetch data:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	// Parse XML data
	var bathInfos BathInfos
	err = xml.Unmarshal(body, &bathInfos)
	if err != nil {
		fmt.Println("Failed to parse XML:", err)
		return
	}

	// Print bath information
	for _, bath := range bathInfos.Baths {
		fmt.Println("Title:", bath.Title)
		fmt.Println("Temperature of water:", bath.TemperatureWater)
		fmt.Println("Status (open/closed):", bath.OpenClosedText)
		fmt.Println("URL:", bath.URLPage)
		fmt.Println()
	}
}
