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
func getAllIps(domains []string) {
	// Prepare request data
	endpoint := fmt.Sprintf("%s/getIps", apiBaseURL)
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

	// Parse response
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %v\n", err)
		return
	}

	// Access ipAddresses map
	if ipData, ok := response["ipAddresses"].(map[string]interface{}); ok {
		// Extract IPv4 addresses
		if ipv4, ok := ipData["ipv4Addresses"].([]interface{}); ok {
			fmt.Println("IPv4 Addresses:")
			for _, ip := range ipv4 {
				if ipStr, ok := ip.(string); ok {
					fmt.Println(ipStr)
				} else {
					fmt.Println("Error: Invalid type in 'ipv4Addresses'")
				}
			}
		} else {
			fmt.Println("Error: 'ipv4Addresses' field not found or not in expected format")
		}

		// Extract IPv6 addresses
		if ipv6, ok := ipData["ipv6Addresses"].([]interface{}); ok {
			fmt.Println("IPv6 Addresses:")
			for _, ip := range ipv6 {
				if ipStr, ok := ip.(string); ok {
					fmt.Println(ipStr)
				} else {
					fmt.Println("Error: Invalid type in 'ipv6Addresses'")
				}
			}
		} else {
			fmt.Println("No 'ipv6Addresses' found or not in expected format")
		}
	} else {
		fmt.Println("Error: 'ipAddresses' field not found or not in expected format")
	}

	// Pretty print the response
	prettyJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}
