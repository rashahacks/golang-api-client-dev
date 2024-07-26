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

func StartCron(cronNotification string, cronTime int64, cronType string) error {
	apiKey := strings.TrimSpace(getAPIKey())
	baseUrl := apiBaseURL
	client := &http.Client{}

	var method = "PUT"
	var url = baseUrl + "/api/v2/startCron"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
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
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(jsonData))

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON response: %v\n", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to start cron: %v\n", response.Message)
	}

	fmt.Println("Message:", response.Message)
	return nil

}

func StopCron() error {
	apiKey := strings.TrimSpace(getAPIKey())
	baseUrl := apiBaseURL
	client := &http.Client{}
	var method = "PUT"
	var url = baseUrl + "/api/v2/stopCron"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("X-Jsmon-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v\n", err)

	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON response: %v\n", err)
	}

	fmt.Println("Message:", response.Message)
	return nil
}

func UpdateCron(cronNotification string, cronType string) error {
	apiKey := strings.TrimSpace(getAPIKey())
	baseUrl := apiBaseURL
	client := &http.Client{}
	var method = "PUT"
	var url = baseUrl + "/api/v2/updateCron"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("X-Jsmon-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"cronJobNotification": cronNotification,
		"vulnerabilitiesType": cronType,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(jsonData))

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v\n", err)
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON response: %v\n", err)
	}

	fmt.Println("Message:", response.Message)
	return nil

}
