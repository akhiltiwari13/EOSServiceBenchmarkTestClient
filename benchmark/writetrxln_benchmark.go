package benchmark

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// MakeWritetrxlnRequest sends request to icury's writetrxln api in 100 go routines
func MakeWritetrxlnRequest(url string, ch chan<- string, id int) {
	start := time.Now()

	// creating payloads for multiple writetx requests
	payload := strings.NewReader(" {\n \"transactionLineId\":" + strconv.Itoa(id) + ",\n \"transactionId\": " + strconv.Itoa(id) + ",\n \"timestamp\": 9999,\n \"icecatUserId\": \"alice\",\n \"value\": \"ICY\",\n \"quantity\": 13,\n \"description\": \"transaction Line description\",\n \"vat\":0\n}")

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
	ch <- fmt.Sprintf("writetrxln test:: %.4f seconds elapsed with response length: %d %s \t id: %d", elapsed, length, url, id)
}
