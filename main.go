package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(argsWithoutProg) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	var maxConcurrency int
	var maxPages int

	if len(argsWithoutProg) > 1 {
		numConc, err := strconv.Atoi(argsWithoutProg[1])
		if err != nil {
			fmt.Println("error parsing maxConcurrency: ", err)
			maxConcurrency = 10
		}
		maxConcurrency = numConc
	} else {
		maxConcurrency = 10
	}

	if len(argsWithoutProg) > 2 {
		numPage, err := strconv.Atoi(argsWithoutProg[2])
		if err != nil {
			fmt.Println("error parsing maxPages: ", err)
			maxPages = 25
		}
		maxPages = numPage
	} else {
		maxPages = 25
	}

	fmt.Println("starting crawl: ", argsWithoutProg[0])
	cfg, err := configure(argsWithoutProg[0], maxConcurrency, maxPages)
	if err != nil {
		fmt.Println("error configure: ", err)
		os.Exit(1)
	}
	cfg.wg.Add(1)
	go cfg.crawlPage(argsWithoutProg[0])
	cfg.wg.Wait()

	PrintReports(cfg.pages, argsWithoutProg[0])
	os.Exit(0)
}
