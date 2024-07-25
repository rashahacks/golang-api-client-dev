package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	scanUrl := flag.String("scanUrl", "", "URL to scan")
	uploadUrl := flag.String("uploadUrl", "", "URL to upload")
	apiKeyFlag := flag.String("apikey", "", "API key for authentication")
	flag.Parse()

	if *apiKeyFlag != "" {
		setAPIKey(*apiKeyFlag)
	} else {
		// loading API key from credentials file
		err := loadAPIKey()
		if err != nil {
			fmt.Println("Error loading API key:", err)
			os.Exit(1)
		}
	}
	fmt.Println("API Key:", getAPIKey())

	// Check if a upload URL is provided
	if *uploadUrl != "" {
		uploadUrlEndpoint(*uploadUrl)
	} else if *scanUrl != "" {
		rescanUrlEndpoint(*scanUrl)
	} else {
		fmt.Println("No action specified. Use --uploadUrl to upload a URL or --scanUrl to rescan a URL.")
	}
}
