package main

import (
	"net/url"
	"strings"
)

func normalizeURL(raw_url string) (string, error) {
	// Remove scheme
	u, err := url.Parse(raw_url)
	if err != nil {
		return "", err
	}
	u.Scheme = ""

	// Remove www
	if u.Hostname()[:4] == "www." {
		u.Host = u.Hostname()[4:]
	}

	// Remove trailing slash
	if len(u.Path) > 1 && u.Path[len(u.Path)-1] == '/' {
		u.Path = u.Path[:len(u.Path)-1]
	}
	url := u.String()

	url = strings.ToLower(url)

	if url[:2] == "//" {
		url = url[2:]
	}

	return url, nil
}
