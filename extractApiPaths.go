package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Function to get API paths based on domains
func getApiPaths(domains []string) {
	// Prepare request data
	endpoint := fmt.Sprintf("%s/apiPathfromDomain", apiBaseURL)
	requestBody, err := json.Marshal(map[string]interface{}{
		"domains": domains,
	})
	if err != nil {
		fmt.Printf("Failed to marshal request body: %v\n", err)
		return
	}

	// Create request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	// fmt.Printf("Sending request to: %s\n", endpoint)
	// fmt.Printf("Request body: %s\n", requestBody)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return
	}

	//fmt.Println("Raw Response Body:")
	//fmt.Println(string(body))

	// Parse response
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %v\n", err)
		return
	}

	// apiPaths, ok := response["apiPaths"].([]interface{})
	// if !ok {
	// 	fmt.Println("Error: 'apiPaths' field not found or not in expected format")
	// 	return
	// }

	// // Print API paths in plain text
	// //fmt.Println("API Paths:")
	// for _, path := range apiPaths {
	// 	if pathStr, ok := path.(string); ok {
	// 		fmt.Println(pathStr)
	// 	} else {
	// 		fmt.Println("Error: Invalid type in 'apiPaths'")
	// 	}
	// }

	prettyJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}