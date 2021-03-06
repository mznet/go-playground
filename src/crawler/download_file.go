package crawler

import (
	"strings"
	"fmt"
	"os"
	"net/http"
	"io"
	"aws"
)

func DownloadFromUrl(url string, timestamp string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)
	path := "./tmp/" + timestamp
	filePath := path + "/" + fileName

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}

	output, err := os.Create(filePath)

	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)

	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
	aws.Aws{}.UploadToS3(filePath)
}
