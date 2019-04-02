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
	"sync"
)

type site struct {
	URL  string `json:"url"`
	UUID string `json:"uuid"`
}

// Concurrent requests
const workersCount = 10

var baseURL = os.Getenv("BASE_URL")

func getURLWorker(siteChan chan map[string]string) {
	for s := range siteChan {
		println(s["url"])
		resp, err := http.Get(s["url"])
		if err != nil {
			log.Fatal(err)
		}

		// Print the HTTP Status Code and Status Name
		fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

		var status = false
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			fmt.Println("HTTP Status is in the 2xx range")
			status = true
		}

		url := baseURL + s["uuid"]

		var jsonStr = []byte(`{"IsHealthy": ` + strconv.FormatBool(status) + `}`)
		req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err = client.Do(req)

		print(resp.Status)
		if err != nil {
			panic(err)
		}

		_ = resp
		_ = err
	}
}

func main() {

	if baseURL == "" {
		panic("BASE_URL env var not configured properly")
	}

	resp, err := http.Get(baseURL)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Printf("%s\n", string(contents))

	var siteMap []map[string]string

	err = json.Unmarshal([]byte(contents), &siteMap)

	var wg sync.WaitGroup
	urlChan := make(chan map[string]string)

	wg.Add(workersCount)

	for i := 0; i < workersCount; i++ {
		go func() {
			getURLWorker(urlChan)
			wg.Done()
		}()
	}

	completed := 0
	for _, each := range siteMap {
		urlChan <- each
		completed++
	}
	close(urlChan)

	wg.Wait()
}
