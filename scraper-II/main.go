package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	blogTitles, err := GetLatestBlogsTitles("http://golangcode.com")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Blog Titles:")
	fmt.Println(blogTitles)
}

// GetLatestBlogsTitles crawls the given url using the net/http
// and uses goquery package to extract latest blogs titles from the resource url points to.
func GetLatestBlogsTitles(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	titles := ""

	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
		titles += "-" + s.Text() + "\n"
	})

	return titles, nil
}
