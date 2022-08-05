package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})

	c.OnHTML(".MRMini", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnHTML("tr > td .body-font-color", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://www.dorics.com/bk_league/907") // Visit 要放最後
}
