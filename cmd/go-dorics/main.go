package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

var host = "https://www.dorics.com"

func getURL(link string) string {
	return host + link
}

func main() {
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println(string(r.Body))
	})

	// c.OnHTML(".MRMini", func(e *colly.HTMLElement) {
	// 	fmt.Println(strings.TrimSpace(e.Text))
	// })

	c.OnHTML("tr > td .body-font-color", func(e *colly.HTMLElement) {
		fmt.Println("Team:", e.Text)
	})

	// (全场)初盘
	c.OnHTML("tbody tr:nth-child(1)", func(e *colly.HTMLElement) {
		e.ForEach("td:nth-child(10)", func(index int, e *colly.HTMLElement) {
			fmt.Println("(全场)初盘:", strings.TrimSpace(e.Text))
		})
	})

	// 实际
	c.OnHTML("tbody tr:nth-child(1)", func(e *colly.HTMLElement) {
		e.ForEach("td:nth-child(11)", func(index int, e *colly.HTMLElement) {
			fmt.Println("实际:", strings.TrimSpace(e.Text))
		})
	})
	// c.OnHTML("tbody tr td.BR0:nth-child(1)", func(e *colly.HTMLElement) {
	// 	fmt.Println("second team:", strings.TrimSpace(e.Text))
	// })

	// On every a element which has href attribute call callback
	c.OnHTML("li a.data-url", func(e *colly.HTMLElement) {
		link := e.Attr("data-url")
		// Print link
		fmt.Println(getURL(link))
		// Visit link found on page on a new thread
		// e.Request.Visit(getURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(host + "/bk_league/383/p.1?type=ended_race") // Visit 要放最後
}
