package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type stringSliceFlag []string

func (s *stringSliceFlag) String() string {
	return strings.Join(*s, ", ")
}

func (s *stringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}
func main() {
	scanUrl := flag.String("scanUrl", "", "URL or scan ID to rescan")
	uploadUrl := flag.String("uploadUrl", "", "URL to upload for scanning")
	apiKeyFlag := flag.String("apikey", "", "API key for authentication")
	scanFileId := flag.String("scanFile", "", "File ID to scan")
	uploadFile := flag.String("uploadFile", "", "File to upload giving path to the file locally.")
	getAllResults := flag.String("automationData", "", "Get all automation results")
	size := flag.Int("size", 10000, "Number of results to fetch (default 10000)")
	getScannerResultsFlag := flag.Bool("scannerData", false, "Get scanner results")
	cron := flag.String("cron", "", "Set cronjob.")
	cronNotification := flag.String("notifications", "", "Set cronjob notification channel.")
	cronTime := flag.Int64("time", 0, "Set cronjob time.")
	cronType := flag.String("vulnerabilitiesType", "", "Set type[URLs, Analysis, Scanner] of cronjob.")
	cronDomains := flag.String("domains", "", "Set domains for cronjob.")
	cronDomainsNotify := flag.String("domainsNotify", "", "Set notify(true/false) for each domain for cronjob.")
	viewurls := flag.Bool("urls", false, "view all urls")
	viewurlsSize := flag.Int("urlSize", 10, "Number of URLs to fetch")
	scanDomainFlag := flag.String("scanDomain", "", "Domain to automate scan")
	wordsFlag := flag.String("words", "", "Comma-separated list of words to include in the scan")

	getDomainsFlag := flag.Bool("getDomains", false, "Get all domains for the user")
	var headers stringSliceFlag
	flag.Var(&headers, "H", "Custom headers in the format 'Key: Value' (can be used multiple times)")

	usageFlag := flag.Bool("usage", false, "View user profile")
	viewfiles := flag.Bool("files", false, "view all files")
	viewEmails := flag.String("Emails", "", "view all Emails for specified domains")
	s3domains := flag.String("S3Domains", "", "get all S3Domains for specified domains")
	ips := flag.String("ips", "", "get all IPs for specified domains")
	domainUrl := flag.String("DomainUrls", "", "get Domain URLs for specified domains")
	apiPath := flag.String("api", "", "get the APIs for specified domains")
	compareFlag := flag.String("compare", "", "Compare two js responses by jsmon_ids (format: JSMON_ID1,JSMON_ID2)")
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s [flags]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Flags:\n")

		fmt.Fprintf(os.Stderr, "INPUT:\n")
		fmt.Fprintf(os.Stderr, "  -scanUrl string         URL or scan ID to rescan\n")
		fmt.Fprintf(os.Stderr, "  -uploadUrl string       URL to upload for scanning\n")
		fmt.Fprintf(os.Stderr, "  -scanFile string        File ID to scan\n")
		fmt.Fprintf(os.Stderr, "  -uploadFile string      File to upload (local path)\n")
		fmt.Fprintf(os.Stderr, "  -scanDomain string      Domain to automate scan\n")

		fmt.Fprintf(os.Stderr, "\nAUTHENTICATION:\n")
		fmt.Fprintf(os.Stderr, "  -apikey string          API key for authentication\n")

		fmt.Fprintf(os.Stderr, "\nOUTPUT:\n")
		fmt.Fprintf(os.Stderr, "  -automationData string  Get all automation results\n")
		fmt.Fprintf(os.Stderr, "  -scannerData            Get scanner results\n")
		fmt.Fprintf(os.Stderr, "  -urls                   View all URLs\n")
		fmt.Fprintf(os.Stderr, "  -size int               Number of URLs to fetch (default 10)\n")
		fmt.Fprintf(os.Stderr, "  -files                  View all files\n")
		fmt.Fprintf(os.Stderr, "  -usage                  View user profile\n")

		fmt.Fprintf(os.Stderr, "\nCRON JOB:\n")
		fmt.Fprintf(os.Stderr, "  -cron string            Set, update, or stop cronjob\n")
		fmt.Fprintf(os.Stderr, "  -notifications string   Set cronjob notification channel\n")
		fmt.Fprintf(os.Stderr, "  -time int               Set cronjob time\n")
		fmt.Fprintf(os.Stderr, "  -vulnerabilitiesType    Set type of cronjob (URLs, Analysis, Scanner)\n")
		fmt.Fprintf(os.Stderr, "  -domains string         Set domains for cronjob\n")
		fmt.Fprintf(os.Stderr, "  -domainsNotify string   Set notify (true/false) for each domain\n")

		fmt.Fprintf(os.Stderr, "\nADDITIONAL OPTIONS:\n")
		fmt.Fprintf(os.Stderr, "  -H string               Custom headers (Key: Value, can be used multiple times)\n")
		fmt.Fprintf(os.Stderr, "  -words string           Comma-separated list of words to include in the scan\n")
		fmt.Fprintf(os.Stderr, "  -getDomains             Get all domains for the user\n")
		fmt.Fprintf(os.Stderr, "  -Emails string          View all Emails for specified domains\n")
		fmt.Fprintf(os.Stderr, "  -S3Domains string       Get all S3 Domains for specified domains\n")
		fmt.Fprintf(os.Stderr, "  -ips string             Get all IPs for specified domains\n")
		fmt.Fprintf(os.Stderr, "  -DomainUrls string      Get Domain URLs for specified domains\n")
		fmt.Fprintf(os.Stderr, "  -api string             Get the APIs for specified domains\n")
		fmt.Fprintf(os.Stderr, "  -compare string         Compare two JS responses by JSMON_IDs (format: ID1,ID2)\n")
	}

	// Handle API key
	if *apiKeyFlag != "" {
		setAPIKey(*apiKeyFlag)
	} else {
		err := loadAPIKey()
		if err != nil {
			fmt.Println("Error loading API key:", err)
			fmt.Println("Please provide an API key using the -apikey flag.")
			os.Exit(1)
		}
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
		uploadFileEndpoint(*uploadFile, headers)
	case *viewurls:
		viewUrls(*viewurlsSize)
	case *viewfiles:
		viewFiles()
	case *uploadUrl != "":
		uploadUrlEndpoint(*uploadUrl, headers)
	case *viewEmails != "":
		domains := strings.Split(*viewEmails, ",")
		for i, domain := range domains {
			domains[i] = strings.TrimSpace(domain)
		}
		getEmails(domains)
	case *s3domains != "":
		domains := strings.Split(*s3domains, ",")
		for i, domain := range domains {
			domains[i] = strings.TrimSpace(domain)
		}
		getS3Domains(domains)
	case *ips != "":
		domains := strings.Split(*ips, ",")
		for i, domain := range domains {
			domains[i] = strings.TrimSpace(domain)
		}
		getAllIps(domains)
	case *domainUrl != "":
		domains := strings.Split(*domainUrl, ",")
		for i, domain := range domains {
			domains[i] = strings.TrimSpace(domain)
		}
		getDomainUrls(domains)
	case *apiPath != "":
		domains := strings.Split(*apiPath, ",")
		for i, domain := range domains {
			domains[i] = strings.TrimSpace(domain)
		}
		getApiPaths(domains)
	case *getScannerResultsFlag:
		getScannerResults()
	case *scanUrl != "":
		rescanUrlEndpoint(*scanUrl)
	case *cron == "start":
		StartCron(*cronNotification, *cronTime, *cronType, *cronDomains, *cronDomainsNotify)
	case *cron == "stop":
		StopCron()
	case *getDomainsFlag:
		getDomains()
	case *compareFlag != "":
		ids := strings.Split(*compareFlag, ",")
		if len(ids) != 2 {
			fmt.Println("Invalid format for compare. Use: JSMON_ID1,JSMON_ID2")
			os.Exit(1)
		}
		compareEndpoint(strings.TrimSpace(ids[0]), strings.TrimSpace(ids[1]))
	case *cron == "update":
		UpdateCron(*cronNotification, *cronType, *cronDomains, *cronDomainsNotify, *cronTime)
	case *getAllResults != "":
		getAllAutomationResults(*getAllResults, *size)
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
