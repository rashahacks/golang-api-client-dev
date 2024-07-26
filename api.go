package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fmt.Printf("File size: %d bytes\n", fileInfo.Size())
	if fileInfo.Size() == 0 {
		fmt.Println("Error: File is empty")
		return
	}

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("First 100 bytes of file content: %s\n", string(fileContent[:min(100, len(fileContent))]))

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file content:", err)
		return
	}

	// Add Content-Type as a form field
	//writer.WriteField("Content-Type", "multipart/form-data; boundary")

	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing multipart writer:", err)
		return
	}

	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	fmt.Printf("Uploading file: %s\n", filePath)
	fmt.Printf("Content-Type: %s\n", writer.FormDataContentType())

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	// Print request details
	fmt.Printf("Request URL: %s\n", req.URL)
	fmt.Printf("Request Method: %s\n", req.Method)
	fmt.Printf("Request Headers:\n")
	for key, values := range req.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", string(responseBody))

	// Parse and print JSON response
	var result interface{}
	err = json.Unmarshal(responseBody, &result)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
