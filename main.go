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
	viewurls := flag.Bool("urls", false, "view all urls")
	scanDomainFlag := flag.String("scanDomain", "", "Domain to automate scan")
	wordsFlag := flag.String("words", "", "Comma-separated list of words to include in the scan")

	usageFlag := flag.Bool("usage", false, "View user profile")
	viewfiles := flag.Bool("files", false, "view all files")
	viewEmails := flag.String("Emails", "", "view all Emails")
	s3domains := flag.String("S3Domains", "", "get all S3Domains")
	ip := flag.String("ips", "", "get all Ips")
	domainUrl := flag.String("DomainUrls", "", "get DomainUrls")
	apiPath := flag.String("api", "", "get the apis")
	compareFlag := flag.String("compare", "", "Compare two js responses by jsmon_ids (format: JSMON_ID1,JSMON_ID2)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
 
	flag.Parse()

	// Handle API key
	if *apiKeyFlag != "" {
		setAPIKey(*apiKeyFlag)
		fmt.Println("Using provided API key.")
	} else {
		err := loadAPIKey()
		if err != nil {
			fmt.Println("Error loading API key:", err)
			fmt.Println("Please provide an API key using the -apikey flag.")
			os.Exit(1)
		}
		fmt.Println("Using API key from credentials file.")
	}

	// Check if any action flags were provided
	if flag.NFlag() == 0 || (flag.NFlag() == 1 && *apiKeyFlag != "") {
		fmt.Println("No action specified. Use -h or --help for usage information.")
		flag.Usage()
		os.Exit(1)
	}

	// Execute the appropriate function based on the provided flag
	switch {
	case *scanFileId != "":
		scanFileEndpoint(*scanFileId)
	case *uploadFile != "":
		uploadFileEndpoint(*uploadFile)
	case *viewurls:
		viewUrls()
	case *viewfiles:
		viewFiles()
	case *uploadUrl != "":
		uploadUrlEndpoint(*uploadUrl)
	case *viewEmails != "":
		// Extract emails for provided domains
		domains := strings.Split(*viewEmails, ",")
		for i, domain := range domains {
			domains[i] = strings.TrimSpace(domain) // Trim any extra spaces
		}
		getEmails(domains)
	case *s3domains !="":
		domains := strings.Split(*s3domains, ",")
		for i, domain := range domains {
			domains[i] = strings.TrimSpace(domain) // Trim any extra spaces
		}
		getS3Domains(domains)
	case *ip !="":
		domains := strings.Split(*ip, ",")
		for i, domain := range domains {
			domains[i] = strings.TrimSpace(domain) // Trim any extra spaces
		}
		getAllIps(domains)
	case *domainUrl !="":
		domains := strings.Split(*domainUrl, ",")
		for i, domain := range domains {
			domains[i] = strings.TrimSpace(domain) // Trim any extra spaces
		}
	getDomainUrls(domains)
	case *apiPath != "": // Split the comma-separated string into a slice
	domains := strings.Split(*apiPath, ",")
	for i, domain := range domains {
		domains[i] = strings.TrimSpace(domain) // Trim any extra spaces
	}
	getApiPaths(domains)
	case *getScannerResultsFlag:
		getScannerResults()
	case *scanUrl != "":
		rescanUrlEndpoint(*scanUrl)
	case *cron == "start":
		StartCron(*cronNotification, *cronTime, *cronType)
	case *cron == "stop":
		StopCron()
	case *compareFlag != "":
		ids := strings.Split(*compareFlag, ",")
		if len(ids) != 2 {
			fmt.Println("Invalid format for compare. Use: JSMON_ID1,JSMON_ID2")
			os.Exit(1)
		}
		compareEndpoint(strings.TrimSpace(ids[0]), strings.TrimSpace(ids[1]))
	case *cron == "update":
		UpdateCron(*cronNotification, *cronType)
	case *getAllResults != "":
		parts := strings.Split(*getAllResults, ",")
		if len(parts) != 3 {
			fmt.Println("Invalid format for automationData. Use: input,inputType,showOnly")
			os.Exit(1)
		}
		getAllAutomationResults(parts[0], parts[1], parts[2])
	case *scanDomainFlag != "":
		words := []string{}
		if *wordsFlag != "" {
			words = strings.Split(*wordsFlag, ",")
		}
		fmt.Printf("Domain: %s, Words: %v\n", *scanDomainFlag, words) // Add this line for debugging
		automateScanDomain(*scanDomainFlag, words)
	case *usageFlag:
		callViewProfile()
	default:
		fmt.Println("No valid action specified.")
		flag.Usage()
		os.Exit(1)
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
