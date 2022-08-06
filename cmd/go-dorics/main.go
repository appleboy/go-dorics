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
	urlList := map[string]struct{}{}
	urlList[getURL("/bk_league/383/p.1?type=ended_race")] = struct{}{}

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
	})

	c.OnHTML(".raceItem table", func(e *colly.HTMLElement) {
		count := 0
		e.ForEach("tr th", func(index int, e *colly.HTMLElement) {
			if strings.TrimSpace(e.Text) == "实际" {
				count = index
			}
		})

		e.ForEach(".body-font-color", func(index int, e *colly.HTMLElement) {
			fmt.Println("Team:", e.Text)
		})

		e.ForEach("tbody tr:nth-child(1)", func(index int, e *colly.HTMLElement) {
			e.ForEach("tbody tr td:nth-child("+fmt.Sprintf("%d", (count+1))+")", func(index int, e *colly.HTMLElement) {
				fmt.Println("(全场)初盘:", strings.TrimSpace(e.Text))
			})

			e.ForEach("tbody tr td:nth-child("+fmt.Sprintf("%d", (count+2))+")", func(index int, e *colly.HTMLElement) {
				fmt.Println("实际:", strings.TrimSpace(e.Text))
			})
		})
	})

	// On every a element which has href attribute call callback
	c.OnHTML("li a.data-url", func(e *colly.HTMLElement) {
		link := getURL(e.Attr("data-url"))
		if _, ok := urlList[link]; !ok {
			urlList[link] = struct{}{}
			// Visit link found on page on a new thread
			e.Request.Visit(link)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(host + "/bk_league/383/p.1?type=ended_race") // Visit 要放最後
}
