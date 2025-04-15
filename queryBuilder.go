package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// QueryBuilder executes a POST request to the /queryBuilder endpoint with the given workspace ID and query.
// It expects the caller to provide the API base URL and API key.
func queryBuilder(wkspId, query string) {
	endpoint := fmt.Sprintf("%s/queryBuilder?wkspId=%s", apiBaseURL, wkspId)

	requestBody, err := json.Marshal(map[string]interface{}{
		"query": query,
	})
	if err != nil {
		fmt.Printf("Failed to marshal request body: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(apiKey))

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %v\n", err)
		return
	}

	
	// Print JS URLs
if urls, ok := response["urls"].([]interface{}); ok {
	fmt.Println("JS URLs:")
	for _, item := range urls {
		if url, ok := item.(string); ok {
			fmt.Println(url)
		}
	}
} else {
	fmt.Println("No JS URLs found in 'urls' field.")
}

// Print paginatedResults if available
if paginatedResults, ok := response["paginatedResults"].([]interface{}); ok && len(paginatedResults) > 0 {
	fmt.Println("\nPaginated Results:")
	switch paginatedResults[0].(type) {
	case string:
		var results []string
		for _, item := range paginatedResults {
			if val, ok := item.(string); ok {
				results = append(results, val)
			}
		}
		output, _ := json.MarshalIndent(results, "", "    ")
		fmt.Println(string(output))
	default:
		output, _ := json.MarshalIndent(paginatedResults, "", "    ")
		fmt.Println(string(output))
	}
} else {
	fmt.Println("No paginated results found.")
}

}
