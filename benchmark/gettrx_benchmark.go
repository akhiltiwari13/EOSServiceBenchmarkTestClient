package benchmark

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// MakeGettrxRequest sends request to icury's gettrx api in 100 go routines
func MakeGettrxRequest(url string, ch chan<- string) {
	start := time.Now()

	req, _ := http.NewRequest("GET", url, nil)

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
	ch <- fmt.Sprintf("gettrx test:: %.4f seconds elapsed with response length: %d %s \t", elapsed, length, url)
}
