package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)


func uploadUrlEndpoint(url string) {
	endpoint := fmt.Sprintf("%s/uploadUrl", apiBaseURL)

	//request body
	requestBody, err := json.Marshal(map[string]string{
		"url": url,
	})
	if err != nil {
		fmt.Println("Error creating request body:", err)
		return
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse and print JSON response
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}

func rescanUrlEndpoint(scanId string) {
	endpoint := fmt.Sprintf("%s/rescanURL/%s", apiBaseURL, scanId)

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse and print JSON response
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}

func scanFileEndpoint(fileId string) {
	endpoint := fmt.Sprintf("%s/scanFile/%s", apiBaseURL, fileId)

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse and print JSON response
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}

func uploadFileEndpoint(filePath string) {
	endpoint := fmt.Sprintf("%s/uploadFile", apiBaseURL)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		log.Fatalf("Error creating form file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatalf("Error copying file content: %v", err)
	}

	err = writer.Close()
	if err != nil {
		log.Fatalf("Error closing multipart writer: %v", err)
	}

	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	log.Printf("Sending request to: %s", endpoint)
	log.Printf("Request headers:")
	for k, v := range req.Header {
		log.Printf("%s: %s", k, v)
	}

	// Debug request body length
	log.Printf("Request body length: %d bytes", body.Len())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	log.Printf("Response status: %s", resp.Status)
	log.Printf("Response body: %s", string(responseBody))

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Upload failed with status code: %d", resp.StatusCode)
	}
}

func getAllAutomationResults(input, inputType string, showOnly string) {
	endpoint := fmt.Sprintf("%s/getAllAutomationResults", apiBaseURL)

	//URL with query parameters
	url := fmt.Sprintf("%s?showonly=%s&inputType=%s&input=%s", endpoint, showOnly, inputType, input)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse and print JSON response
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}

func getScannerResults() {
	endpoint := fmt.Sprintf("%s/getScannerResults", apiBaseURL)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse JSON response
	var result struct {
		Message string `json:"message"`
		Data    struct {
			ModuleName []string `json:"moduleName"`
			URL        string   `json:"url"`
		} `json:"data"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print the result in a more readable format
	fmt.Println("Message:", result.Message)
	fmt.Println("URL:", result.Data.URL)
	fmt.Println("Modules found:")
	for _, module := range result.Data.ModuleName {
		fmt.Printf("- %s\n", module)
	}
}

