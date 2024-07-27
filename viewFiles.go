package main
import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type FilesResponse struct {
	Files   []string `json:"files"`
	Message string   `json:"message"`
}

func viewFiles() {
	endpoint := fmt.Sprintf("%s/viewFiles", apiBaseURL)
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Errorf("failed to create request: %v", err)
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

	var response FilesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("failed to unmarshal JSON response: %v", err)
		return
	}

	fmt.Println("Message:", response.Message)
	fmt.Println("Files:", response.Files)
}