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

	scanUrl := flag.String("scanUrl", "", "URL or scan ID to rescan")
	uploadUrl := flag.String("uploadUrl", "", "URL to upload for scanning")
	apiKeyFlag := flag.String("apikey", "", "API key for authentication")
	scanFileId := flag.String("scanFile", "", "File ID to scan")
	uploadFile := flag.String("uploadFile", "", "File to upload giving path to the file locally.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Check if any flags were provided
	if flag.NFlag() == 0 {
		fmt.Println("No options provided. Use -h or --help for usage information.")
		flag.Usage()
		os.Exit(1)
	}
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

	if *scanFileId != "" {
		scanFileEndpoint(*scanFileId)
	} else if *uploadFile != "" {
		uploadFileEndpoint(*uploadFile)
	} else if *uploadUrl != "" {
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
		flag.Usage()
	}
}
