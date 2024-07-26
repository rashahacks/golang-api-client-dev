package main

import (
	"flag"
	"fmt"
	"os"
)

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

func main() {
	scanUrl := flag.String("scanUrl", "", "URL to scan")
	uploadUrl := flag.String("uploadUrl", "", "URL to upload")
	apiKeyFlag := flag.String("apikey", "", "API key for authentication")
	cron := flag.String("cron", "", "Set cronjob.")
	cronNotification := flag.String("notifications", "", "Set cronjob notification.")
	cronTime := flag.Int64("time", 0, "Set cronjob time.")
	cronType := flag.String("type", "", "Set type of cronjob.")

	flag.Parse()

	//args := parseArgs()

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
	} else if *cron == "start" {
		StartCron(*cronNotification, *cronTime, *cronType)
	} else if *cron == "stop" {
		StopCron()
	} else if *cron == "update" {
		UpdateCron(*cronNotification, *cronType)
	} else {
		fmt.Println("No action specified. Use --uploadUrl to upload a URL or --scanUrl to rescan a URL.")
	}
}
