package main

import (
	"fmt"
	"os"
	"time"
	"crawler"
	"tools"
	"strings"
	"regexp"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"aws"
)

type Config struct {
	AccessKey string `yaml:"access_key"`
	Secret string `yaml:"secret"`
	Token string `yaml: "token"`
	Region string `yaml: "region"`
	Bucket string `yaml: "token"`
}

func main() {
	var imageUrl string
	var urlPattern = regexp.MustCompile(`(https|http):\/\/`)
	var urlCount int = 0

	var config Config
	filepath := ".aws/credential.yaml"

	source, readErr := ioutil.ReadFile(filepath)

	if readErr != nil {
		panic(readErr)
	}

	readErr = yaml.Unmarshal(source, &config)
	if readErr != nil {
		panic(readErr)
	}
	fmt.Printf("Value: %#v\n", config)



	return

	foundUrls := make(map[string]bool)
	seedUrls := os.Args[1:]
	extArray := [2]string { "jpg", "png" }

	timestamp := time.Now().Format("20060102150405")

	// Channels
	chUrls := make(chan string)
	chFinished := make(chan bool)

	// Kick off the crawl process (concurrently)
	for _, url := range seedUrls {
		if urlPattern.MatchString(url) == false {
			fmt.Println(url, "is not acceptable URL format")
			continue
		}
		go crawler.Crawl(url, chUrls, chFinished)
		urlCount += 1
	}

	if urlCount == 0 {
		panic("No URL is provided")
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

		imageUrl = url

		if strings.Contains(url, "http:") {
			imageUrl = url
		}

		fmt.Println(" - " + imageUrl)

		crawler.DownloadFromUrl(imageUrl, timestamp)
	}

	close(chUrls)
}