package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// takes a html token and search for href attribute in the token
// if the token has href attribute returns true,href value
// else the token does not have token return false and ""
func getHref(t html.Token) (ok bool, href string) {
	for _, attr := range t.Attr {
		if attr.Key == "href" {
			href = attr.Val
			ok = true
		}
	}
	return ok, href
}

// takes an url and crawls the resource, the url points to, and
// collects all the anchor tags in the resource and
// return the href value of all the anchor tags to the chanUrls channel
// after completion returns true to the chanFinished channel
func crawl(url string, chanUrls chan string, chanFinished chan bool) {
	resp, err := http.Get(url)
	defer func() {
		chanFinished <- true
	}()
	if err != nil {
		log.Fatalf("ERROR: Failed to connect to the url: %s", url)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("ERROR: Failed to close url: %s", url)
		}
	}(resp.Body)

	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return

		case tt == html.StartTagToken:
			t := z.Token()
			isAnchorTag := t.Data == "a"
			if !isAnchorTag {
				continue
			}
			ok, url := getHref(t)
			if !ok {
				continue
			}

			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				chanUrls <- url
			}
		}
	}
}

func main() {
	foundUrls := make(map[string]bool)
	seedUrls := os.Args[1:]

	chanUrls := make(chan string)
	chanFinished := make(chan bool)

	// Deploy a go-routine for every url user inputs
	for _, url := range seedUrls {
		go crawl(url, chanUrls, chanFinished)
	}

	// wait for the go-routine(s) to find as many urls it can find
	// and wait for all the go-routine(s) until all of them are finished
	for c := 0; c < len(seedUrls); {
		select {
		case url := <-chanUrls:
			foundUrls[url] = true
		case <-chanFinished:
			c++
		}
	}

	// Print the found urls
	fmt.Println("\nFound", len(foundUrls), "unique urls")
	for url := range foundUrls {
		fmt.Println("-" + url)
	}
	close(chanUrls)
	close(chanFinished)
}
