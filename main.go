package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getBody(resp *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String()
}

func poll(url string, interval time.Duration) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("- polling", url)
	defer resp.Body.Close()
	fmt.Printf("+ response %s from %s; body:\n%s\n", resp.Status, url, getBody(resp))
	time.Sleep(interval * time.Millisecond)
}

func main() {
	intervalFlag := flag.Int("i", 3000, "interval between polls (in milliseconds)")
	numberFlag := flag.Int("c", 10, "number of polls to perform (0 for infinite)")
	urlFlag := flag.String("u", "", "url to poll")
	flag.Parse()

	if *intervalFlag <= 0 {
		panic("timeout must be greater than 0")
	}

	if *urlFlag == "" {
		panic("url must be provided")
	}

	if *numberFlag > 0 {
		fmt.Println("\npolling", *urlFlag, "every", *intervalFlag, "milliseconds", *numberFlag, "times\n")
		for i := 0; i < *numberFlag; i++ {
			poll(*urlFlag, time.Duration(*intervalFlag))
		}
		return
	} else {
		fmt.Println("polling", *urlFlag, "every", *intervalFlag, "milliseconds")
		poll(*urlFlag, time.Duration(*intervalFlag))
	}

	fmt.Println("done")
}
