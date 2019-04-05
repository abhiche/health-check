package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/abhiche/health-check/pkg/waitgroup"
)

type site struct {
	URL  string `json:"url"`
	UUID string `json:"uuid"`
}

// Concurrent requests
const workersCount = 10

var healthCheckTimeout = 800 * time.Millisecond

var baseURL = os.Getenv("BASE_URL")

func getURLWorker(s map[string]string) {
	println("\n", s["url"])
	timeout := time.Duration(healthCheckTimeout)
	client := http.Client{
		Timeout: timeout,
	}
	status := false
	resp, err := client.Get(s["url"])
	if err != nil {
		fmt.Printf("\n%s", err)
	} else {
		status = getStatus(resp)
	}

	fmt.Print("status ", strconv.FormatBool(status))

	updateSite(status, s)
}

func getStatus(resp *http.Response) bool {
	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	var status = false
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")
		status = true
	}
	return status
}

func updateSite(status bool, s map[string]string) {

	url := baseURL + s["uuid"]

	var jsonStr = []byte(`{"isHealthy": ` + strconv.FormatBool(status) + `}`)
	fmt.Printf(s["uuid"], ` :  {"isHealthy": `+strconv.FormatBool(status)+`}`)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Print("\nUpdate site failed", s["uuid"], err)
	} else {
		print(resp.Status)
	}

}

// FetchAllSites fetch all the sites stored in the application
func FetchAllSites() ([]map[string]string, error) {
	resp, err := http.Get(baseURL)

	if err != nil {
		fmt.Printf("\n%s", err)
		return nil, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\n%s", err)
		return nil, err
	}
	fmt.Printf("\n%s\n", string(contents))

	var siteMap []map[string]string

	err = json.Unmarshal([]byte(contents), &siteMap)

	return siteMap, nil
}

func main() {

	if baseURL == "" {
		panic("BASE_URL env var not configured properly")
	}

	siteMap, err := FetchAllSites()
	if err != nil {
		log.Fatal("Fetching all the sites failed", err)
		return
	}

	wg := waitgroup.NewWaitGroup(workersCount)

	for _, s := range siteMap {
		wg.BlockAdd()
		println("Adding wait group")
		go func(s map[string]string) {
			println("Calling URL worker")
			getURLWorker(s)
			wg.Done()
		}(s)
	}

	wg.Wait()
}
