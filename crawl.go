package main

import (
	"flag"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

// Define a Set to store Links
var linksSet = make(map[string]struct{})

// Store visited links to visit url just once
var visited = make(map[string]bool)

func main() {
	// Define a command-line flag for the URL
	urlPtr := flag.String("url", "", "URL to crawl")
	depthPtr := flag.Int("depth", 1, "Crawl depth")
	flag.Parse()

	// Check if the URL was provided
	if *urlPtr == "" {
		fmt.Println("Please provide a URL using the -url flag.")
		return
	}

	// Start crawling the provided URL
	crawl(*urlPtr, *depthPtr)
	printSortedLinks()
}

func crawl(url string, depth int) {
	if depth <= 0 || visited[url] {
		return
	}
	visited[url] = true

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t := z.Token()
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						link := a.Val
						if strings.HasPrefix(link, "/") {
							link = url + link
						}
						if strings.HasPrefix(link, "http") {
							// Add link to Set
							linksSet[link] = struct{}{}
							crawl(link, depth-1)
						}
					}
				}
			}
		}
	}
}

func printSortedLinks() {
	var links []string
	for link := range linksSet {
		links = append(links, link)
	}
	sort.Strings(links)
	for _, link := range links {
		fmt.Println(link)
	}
}
