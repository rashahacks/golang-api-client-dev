package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)


func queryBuilder(wkspId, query string) {
	endpoint := fmt.Sprintf("%s/queryBuilder?wkspId=%s", apiBaseURL, wkspId)

	requestBody, err := json.Marshal(map[string]interface{}{
		"query": query,
	})
	if err != nil {
		fmt.Printf("Failed to marshal request body: %v\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("[ERR] Wrong API key")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		os.Exit(1)
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %v\n", err)
		os.Exit(1)
	}

	paginatedResults, ok := response["paginatedResults"].([]interface{})
	if !ok || len(paginatedResults) == 0 {
		fmt.Println("No paginated results found.")
		return
	}

	// Check if it's a list of strings (emails) or structured objects
	switch paginatedResults[0].(type) {
	case string:
		// Just print the emails as an indented array
		var emails []string
		for _, item := range paginatedResults {
			if email, ok := item.(string); ok {
				emails = append(emails, email)
			}
		}
		output, err := json.MarshalIndent(emails, "", "    ")
		if err != nil {
			fmt.Printf("Failed to marshal emails: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(output))

	case map[string]interface{}:
		// Original structured format
		var formattedResults []map[string]interface{}
		for _, result := range paginatedResults {
			if resultMap, ok := result.(map[string]interface{}); ok {
				formattedResult := make(map[string]interface{})

				if url, ok := resultMap["url"].(string); ok {
					formattedResult["url"] = url
				}
				if domainName, ok := resultMap["domainName"].(string); ok {
					formattedResult["domainName"] = domainName
				}

				if detectedWords, ok := resultMap["detectedWords"].([]interface{}); ok {
					var formattedDetectedWords []map[string]interface{}
					for _, item := range detectedWords {
						if wordMap, ok := item.(map[string]interface{}); ok {
							formattedWord := make(map[string]interface{})
							if name, ok := wordMap["name"].(string); ok {
								formattedWord["name"] = name
							}
							if words, ok := wordMap["words"].([]interface{}); ok {
								var wordList []string
								for _, word := range words {
									if wordStr, ok := word.(string); ok {
										wordList = append(wordList, wordStr)
									}
								}
								formattedWord["words"] = wordList
							}
							formattedDetectedWords = append(formattedDetectedWords, formattedWord)
						}
					}
					formattedResult["detectedWords"] = formattedDetectedWords
				}

				formattedResults = append(formattedResults, formattedResult)
			}
		}
		output, err := json.MarshalIndent(formattedResults, "", "    ")
		if err != nil {
			fmt.Printf("Failed to marshal formatted results: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(output))

	default:
		fmt.Println("Unknown result format.")
	}
}
