package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", err
	}
	if !strings.Contains(resp.Header.Get("content-type"), "text/html") {
		return "", errors.New("invalid content type : " + resp.Header.Get("content-type"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Function to check if two URLs share the same domain
func isSameDomain(url1 string, url2 string) (bool, error) {
	parsedURL1, err := url.Parse(url1)
	if err != nil {
		return false, errors.New("error parsing URL1: " + err.Error())
	}

	parsedURL2, err := url.Parse(url2)
	if err != nil {
		return false, errors.New("error parsing URL2: " + err.Error())
	}
	// subdomains are part of the same domain, and should return true
	if parsedURL1.Hostname() == parsedURL2.Hostname() {
		return true, nil
	}

	parts1 := strings.Split(parsedURL1.Hostname(), ".")
	parts2 := strings.Split(parsedURL2.Hostname(), ".")
	mainDomain1 := parsedURL1.Hostname()
	mainDomain2 := parsedURL2.Hostname()
	if len(parts1) >= 2 {
		mainDomain1 = strings.Join(parts1[len(parts1)-2:], ".")
	}
	if len(parts2) >= 2 {
		mainDomain2 = strings.Join(parts2[len(parts2)-2:], ".")
	}
	if mainDomain1 == mainDomain2 {
		return true, nil
	}

	return false, nil
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.checkMaxPages() {
		return
	}

	same, err := isSameDomain(cfg.baseURL.String(), rawCurrentURL)
	if err != nil {
		fmt.Println("error Domain: ", err)
		return
	}

	if !same {
		//fmt.Println("!same domain: ", rawBaseURL, rawCurrentURL)
		return
	}
	normal_url, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("error normalizeURL: ", err)
		return
	}

	isFirst := cfg.addPageVisit(normal_url)
	if !isFirst {
		return
	}
	fmt.Println("looking at: ", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("error getHTML: ", err)
		return
	}
	urls, err := getURLsFromHTML(html, rawCurrentURL)
	if err != nil {
		fmt.Println("error getURLsFromHTML: ", err)
		return
	}
	//fmt.Println("found urls: ", urls)
	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}

}
