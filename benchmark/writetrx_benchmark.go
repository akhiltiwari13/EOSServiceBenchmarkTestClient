package benchmark

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// MakeWritetrxRequest sends request to icury's writetrx api in 100 go routines
func MakeWritetrxRequest(url string, ch chan<- string, transactionID int) {
	start := time.Now()

	// creating payloads for multiple writetx requests
	payload := strings.NewReader("{\n\t\"transactionId\":" + strconv.Itoa(transactionID) + ",\n\t\"icecatUserIdFrom\": \"alice\",\n\t\"icecatUserIdTo\": \"bob\",\n\t\"currency\":\"USD\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("blockchain", "EOS")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "770d96a4-61ec-4eae-97a8-e26ea1a39d56")

	res, err := http.DefaultClient.Do(req)

	var length int
	if err == nil {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(body))
		}
		length = len(body)
	}

	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("writetrx test:: %.4f seconds elapsed with response length: %d %s \t id: %d", elapsed, length, url, transactionID)
}
