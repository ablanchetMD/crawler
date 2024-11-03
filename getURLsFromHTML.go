package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func isAbsolute(link string) bool {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return false
	}
	return parsedURL.IsAbs() // returns true if the URL has a scheme
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	return_list := []string{}
	if err != nil {
		return nil, err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if isAbsolute(a.Val) {
						return_list = append(return_list, a.Val)
					} else {
						baseURL, err := url.Parse(rawBaseURL)
						if err != nil {
							return
						}
						return_list = append(return_list, baseURL.ResolveReference(&url.URL{Path: a.Val}).String())
						break
					}
				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return return_list, nil
}
