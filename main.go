package main

import (
	"fmt"
	"icury_bencmark_tests/benchmark"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	// Testing API writetrx
	fmt.Printf("Testing MakeWritetrxRequest API.......\n")
	url := "http://51.254.99.43:5000/transaction/writetrx"
	for i := 0; i < 100; i++ {
		go benchmark.MakeWritetrxRequest(url, ch, i)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)

	}
	elapsed := time.Since(start).Seconds()
	log.Printf("Total writetrx hits took %f time in seconds", elapsed)

	// Testing API writetrxln
	fmt.Printf("Testing MakeWritetrxlnRequest API.......\n")
	url = "http://51.254.99.43:5000/transaction/writetrxln"
	for i := 0; i < 100; i++ {
		go benchmark.MakeWritetrxlnRequest(url, ch, i)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
	elapsed = time.Since(start).Seconds()
	log.Printf("Total writetrxln hits took %f time in seconds", elapsed)

	// Testing API gettrx
	fmt.Printf("Testing MakeGettrxRequest API.......\n")
	url = "http://51.254.99.43:5000/transaction/gettrx"
	for i := 0; i < 100; i++ {
		go benchmark.MakeGettrxRequest(url, ch)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
	elapsed = time.Since(start).Seconds()
	log.Printf("Total gettrx hits took %f time in seconds", elapsed)

	// Testing API gettrxln
	fmt.Printf("Testing MakeGettrxlnRequest API.......\n")
	url = "http://51.254.99.43:5000/transaction/gettrxln"
	for i := 0; i < 100; i++ {
		go benchmark.MakeGettrxlnRequest(url, ch)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
	elapsed = time.Since(start).Seconds()
	log.Printf("Total gettrxln hits took %f time in seconds", elapsed)

	elapsed = time.Since(start).Seconds()
	log.Printf("Total API hits took %f time in seconds", elapsed)

}
