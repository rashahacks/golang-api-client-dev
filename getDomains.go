package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	// "time"
)

func getDomains() {
	endpoint := fmt.Sprintf("%s/getDomains", apiBaseURL)

	req, err := http.NewRequest("GET", endpoint, nil)
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

	var domains []string
	err = json.Unmarshal(body, &domains)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print each domain on a new line
	for _, domain := range domains {
		fmt.Println(domain)
	}
}