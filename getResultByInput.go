package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getAutomationResultsByInput(inputType, value string, wkspId string) {
	endpoint := fmt.Sprintf("%s/getAllJsUrlsResults?inputType=%s&input=%s&wkspId=%s", apiBaseURL, inputType, value, wkspId)

	// Create a new HTTP request with the GET method
	req, err := http.NewRequest("POST", endpoint, nil) // No need for request body in GET
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		return
	}

	// Set necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey())) // Trim any whitespace from the API key

	// Create an HTTP client and make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("[ERR] Wrong API key")
		return 
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response: %v\n", err)
		return
	}

	// Check if the response is successful (Status Code: 2xx)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// Print the response with all fields related to the jsmonId
		fmt.Println(string(body))
	} else {
		fmt.Printf("Error: Received status code %d\n", resp.StatusCode)
		fmt.Println("Response:", string(body)) // Print the response even if it's an error
	}

}
