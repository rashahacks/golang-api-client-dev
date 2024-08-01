package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func StartCron(cronNotification string, cronTime int64, cronType string) {
	apiKey := strings.TrimSpace(getAPIKey())
	baseUrl := apiBaseURL
	client := &http.Client{}

	var method = "PUT"
	var url = baseUrl + "/startCron"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Errorf("failed to create request: %v", err)
		return
	}
	req.Header.Set("X-Jsmon-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"cronJobNotification": cronNotification,
		"vulnerabilitiesType": cronType,
		"time":                cronTime,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("failed to marshal JSON: %v", err)
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(jsonData))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("failed to send request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("failed to read response body: %v", err)
		return
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Errorf("failed to unmarshal JSON response: %v", err)
		return
	}

	fmt.Println("Message:", response.Message)

}

func StopCron() {
	apiKey := strings.TrimSpace(getAPIKey())
	baseUrl := apiBaseURL
	client := &http.Client{}
	var method = "PUT"
	var url = baseUrl + "/stopCron"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Errorf("failed to create request: %v", err)
		return
	}
	req.Header.Set("X-Jsmon-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("failed to send request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("failed to read response body: %v", err)
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Errorf("failed to unmarshal JSON response: %v", err)
		return
	}

	fmt.Println("Message:", response.Message)

}

func UpdateCron(cronNotification string, cronType string) {
	apiKey := strings.TrimSpace(getAPIKey())
	baseUrl := apiBaseURL
	client := &http.Client{}
	var method = "PUT"
	var url = baseUrl + "/updateCron"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Errorf("failed to create request: %v", err)
		return
	}
	req.Header.Set("X-Jsmon-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"cronJobNotification": cronNotification,
		"vulnerabilitiesType": cronType,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("failed to marshal JSON: %v", err)
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(jsonData))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("failed to send request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("failed to read response body: %v", err)
		return
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Errorf("failed to unmarshal JSON response: %v", err)
		return
	}

	fmt.Println("Message:", response.Message)

}
