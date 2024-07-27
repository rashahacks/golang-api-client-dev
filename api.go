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

type URLResponse struct {
	Urls    []URLItem `json:"urls"`
	Message string    `json:"Message"`
}
type URLItem struct {
	URL string `json:"url"`
}

type AutomateScanDomainRequest struct {
	Domain string `json:"domain"`
}

type AnalysisResult struct {
	Message     string `json:"message"`
	TotalChunks int    `json:"totalChunks"`
}

type ModuleScanResult struct {
	Message string `json:"message"`
	Data    []struct {
		ModuleName string `json:"ModuleName"`
		URL        string `json:"URL"`
	} `json:"data"`
}

type ScanResponse struct {
	Message          string           `json:"message"`
	AnalysisResult   AnalysisResult   `json:"analysis_result"`
	ModuleScanResult ModuleScanResult `json:"modulescan_result"`
}

type AutomateScanDomainResponse struct {
	Message       string       `json:"message"`
	FileId        string       `json:"fileId"`
	TrimmedDomain string       `json:"trimmedDomain"`
	ScanResponse  ScanResponse `json:"scanResponse"`
}

func uploadUrlEndpoint(url string) {
	endpoint := fmt.Sprintf("%s/uploadUrl", apiBaseURL)

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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

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

	fmt.Println("Message:", result.Message)
	fmt.Println("URL:", result.Data.URL)
	fmt.Println("Modules found:")
	for _, module := range result.Data.ModuleName {
		fmt.Printf("- %s\n", module)
	}
}

func viewUrls() {
	fmt.Println("viewUrls function called")
	endpoint := fmt.Sprintf("%s/searchAllUrls", apiBaseURL)
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("failed to send request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v", err)
		return
	}
	var response URLResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("failed to unmarshal JSON response: %v", err)
		return
	}

	fmt.Println("Message:", response.Message)
	fmt.Println("URLs:", response.Urls)
}

func automateScanDomain(domain string) {
	fmt.Println("automateScanDomain function called")
	endpoint := fmt.Sprintf("%s/automateScanDomain", apiBaseURL)

	requestBody := AutomateScanDomainRequest{Domain: domain}
	body, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("failed to marshal request body: %v\n", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("failed to send request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("non-200 response: %s\n", responseBody)
		return
	}

	var response AutomateScanDomainResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Printf("failed to unmarshal JSON response: %v\n", err)
		return
	}

	printFormattedResponse(response)
}

func printFormattedResponse(response AutomateScanDomainResponse) {
	fmt.Println("Message:", response.Message)
	fmt.Println("File ID:", response.FileId)
	fmt.Println("Trimmed Domain:", response.TrimmedDomain)

	fmt.Println("\nScan Response:")
	fmt.Println("  Message:", response.ScanResponse.Message)

	fmt.Println("\n  Analysis Result:")
	fmt.Println("    Message:", response.ScanResponse.AnalysisResult.Message)
	fmt.Println("    Total Chunks:", response.ScanResponse.AnalysisResult.TotalChunks)

	fmt.Println("\n  Module Scan Result:")
	fmt.Println("    Message:", response.ScanResponse.ModuleScanResult.Message)
	for _, module := range response.ScanResponse.ModuleScanResult.Data {
		fmt.Println("    Module Name:", module.ModuleName)
		fmt.Println("    URL:", module.URL)
		fmt.Println()
	}
}

func callViewProfile() {
	endpoint := fmt.Sprintf("%s/viewProfile", apiBaseURL)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Jsmon-Key", strings.TrimSpace(getAPIKey()))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		os.Exit(1)
	}

	if data, ok := result["data"].(map[string]interface{}); ok {
		var accountType string
		if orgFound, ok := data["orgFound"].(bool); ok && orgFound {
			accountType = "org"
		} else if personalProfile, ok := data["personalProfile"].(bool); ok && personalProfile {
			accountType = "user"
		} else {
			accountType = "unknown"
		}

		filteredResult := map[string]interface{}{
			"limits": data["apiCallLimits"],
			"type":   accountType,
		}
		filteredData, _ := json.MarshalIndent(filteredResult, "", "  ")
		fmt.Println(string(filteredData))
	} else {
		fmt.Println("Error: Invalid response format")
	}
}

// getAllAutomationResults - > --AutomationData (flag name) with showonly to be changed as View and no sort and pagination
// input -> inputValue other options are compulsory for this function

// getScannerResult -> --scannerData (flag Name)
