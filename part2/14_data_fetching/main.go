package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type IPResponse struct {
    IP string `json:"ip"`
}

type City struct {
	Name string `json:"name"`
}

type Country struct {
	Name string `json:"name"`
}

type Location struct {
	Country Country `json:"country"`
	City    City `json:"city"`
}

type GeoResponse struct {
	Data struct {
		Location Location `json:"location"`
	} `json:"data"`
}

func main() {
	// Get IP
    url1 := "https://api.ipify.org?format=json"

    response, err := http.Get(url1)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        fmt.Println("Error: HTTP Status Code", response.StatusCode)
        return
    }

    var ipResponse IPResponse
    decoder := json.NewDecoder(response.Body)
    if err := decoder.Decode(&ipResponse); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    fmt.Println("IPv4 Address:", ipResponse.IP)

	// Get Geo Data from IP
	// if the token expires get another one at https://ipbase.com/products/ip-api/ from the sandbox
	apiKey := "sgiPfh4j3aXFR3l2CnjWqdKQzxpqGn9pX5b3CUsz"
	url2 := "https://api.ipbase.com/v2/info?apikey=" + apiKey + "&ip=" + ipResponse.IP
	response2, err := http.Get(url2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response2.Body.Close()


    if response2.StatusCode != http.StatusOK {
        fmt.Println("Error: HTTP Status Code", response2.StatusCode)
        return
    }

    var geoData GeoResponse
    decoder = json.NewDecoder(response2.Body)
    if err = decoder.Decode(&geoData); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }
    fmt.Println("Country:", geoData.Data.Location.Country.Name)
    fmt.Println("City:", geoData.Data.Location.City.Name)
}
