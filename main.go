package main

import (
	"fmt"
	"os"
	"time"
	"crawler"
	"tools"
	"strings"
)

type Response struct {
	hostname string
	path string
}

func main() {
	var imageUrl string

	foundUrls := make(map[string]bool)
	seedUrls := os.Args[1:]
	extArray := [2]string { "jpg", "png" }

	timestamp := time.Now().Format("20060102150405")

	// Channels
	chUrls := make(chan string)
	chFinished := make(chan bool)

	// Kick off the crawl process (concurrently)
	for _, url := range seedUrls {
		go crawler.Crawl(url, chUrls, chFinished)
	}

	// Subscribe to both channels
	for c := 0; c < len(seedUrls); {
		select {
		case url := <-chUrls:
			foundUrls[url] = true
		case <-chFinished:
			c++
		}
	}


	fmt.Println("\nFound", len(foundUrls), "unique image source:\n")

	for url, _ := range foundUrls {
		splitted := strings.Split(url, ".")
		splittedCount := len(splitted)

		if tools.StringInclude(extArray, splitted[splittedCount - 1]) != true {
			continue

		}

		imageUrl = "http:" + url

		if strings.Contains(url, "http:") {
			imageUrl = url
		}

		fmt.Println(" - " + imageUrl)

		crawler.DownloadFromUrl(imageUrl, timestamp)
	}

	close(chUrls)
}