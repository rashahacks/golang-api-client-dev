package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanUrl := flag.String("scanUrl", "", "URL or scan ID to rescan")
	uploadUrl := flag.String("uploadUrl", "", "URL to upload for scanning")
	apiKeyFlag := flag.String("apikey", "", "API key for authentication")
	scanFileId := flag.String("scanFile", "", "File ID to scan")
	uploadFile := flag.String("uploadFile", "", "File to upload giving path to the file locally.")
	getAllResults := flag.String("automationData", "", "Get all automation results")
	getScannerResultsFlag := flag.Bool("scannerData", false, "Get scanner results")
	cron := flag.String("cron", "", "Set cronjob.")
	cronNotification := flag.String("notifications", "", "Set cronjob notification.")
	cronTime := flag.Int64("time", 0, "Set cronjob time.")
	cronType := flag.String("type", "", "Set type of cronjob.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	// Check if any flags were provided
	if flag.NFlag() == 0 {
		fmt.Println("No options provided. Use -h or --help for usage information.")
		flag.Usage()
		os.Exit(1)
	}

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

	if *scanFileId != "" {
		scanFileEndpoint(*scanFileId)
	} else if *uploadFile != "" {
		uploadFileEndpoint(*uploadFile)
	} else if *uploadUrl != "" {
		uploadUrlEndpoint(*uploadUrl)
	} else if *getScannerResultsFlag {
		getScannerResults()
	} else if *scanUrl != "" {
		rescanUrlEndpoint(*scanUrl)
	} else if *cron == "start" {
		StartCron(*cronNotification, *cronTime, *cronType)
	} else if *cron == "stop" {
		StopCron()
	} else if *cron == "update" {
		UpdateCron(*cronNotification, *cronType)
	} else if *getAllResults != "" {
		parts := strings.Split(*getAllResults, ",")
		if len(parts) != 3 {
			fmt.Println("Invalid format for getAllResults. Use: input,inputType,showOnly")
			return
		}
		getAllAutomationResults(parts[0], parts[1], parts[2])
	} else {
		fmt.Println("No action specified.")
		flag.Usage()
	}
}

// type Args struct {
// 	Cron             string
// 	CronNotification string
// 	CronTime         int64
// 	CronType         string
// }

// func parseArgs() Args {
// 	//CRON JOB FLAGS ->
// 	cron := flag.String("cron", "", "Set cronjob.")
// 	cronNotification := flag.String("notifications", "", "Set cronjob notification.")
// 	cronTime := flag.Int64("time", 0, "Set cronjob time.")
// 	cronType := flag.String("type", "", "Set type of cronjob.")

// 	flag.Parse()

// 	return Args{
// 		Cron:             *cron,
// 		CronNotification: *cronNotification,
// 		CronTime:         *cronTime,
// 		CronType:         *cronType,
// 	}
// }
