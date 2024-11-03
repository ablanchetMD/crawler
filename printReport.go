package main

import (	
	"fmt"
	"sort"
)

type page struct {
	URL   string
	Count int
}

func PrintReports(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Println("  REPORT for " + baseURL)
	fmt.Println("=============================")

	var sortedPages []page

	for p, count := range pages {
		sortedPages = append(sortedPages, page{URL: p, Count: count})
	}

	sort.Slice(sortedPages, func(i, j int) bool {
		if sortedPages[i].Count == sortedPages[j].Count {
			return sortedPages[i].URL < sortedPages[j].URL
		}

		return sortedPages[i].Count > sortedPages[j].Count
	})

	for _, p := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", p.Count, p.URL)
	}
}