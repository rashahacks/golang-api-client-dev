package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getDomainUrls(domains []string) {
	endpoint := fmt.Sprintf("%s/getDomainsUrls", apiBaseURL)
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

	// Print raw response for debugging
	// fmt.Println("Raw Response Body:")
	// fmt.Println(string(body))

	// Parse response
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %v\n", err)
		return
	}

	getUrls, ok := response["getUrls"].([]interface{})
	if !ok {
		fmt.Println("Error: 'getUrls' field not found or not in expected format")
		return
	}

	// Print URLs in plain text
	for _, url := range getUrls {
		if urlStr, ok := url.(string); ok {
			fmt.Println(urlStr)
		} else {
			fmt.Println("Error: Invalid type in 'getUrls'")
		}
	}
}
