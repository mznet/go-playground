package crawler

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
)

// Extract all http** links from a given webpage
func Crawl(url string, ch chan string, chFinished chan bool) {
	resp, err := http.Get(url)

	defer func() {
		// Notify that we're done after this function
		chFinished <- true
	}()

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}

	b := resp.Body

	defer b.Close() // close Body when the function returns

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			// End of the document, we're done
			return
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "img"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, url := GetHref(t)
			if !ok {
				continue
			}

			ch <- url
		}
	}
}