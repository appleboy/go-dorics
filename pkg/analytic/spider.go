package analytic

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

var host = "https://www.dorics.com"

func getURL(link string) string {
	return host + link
}

func Spider(bk string) *Board {
	cueerntURL := ""
	board := &Board{
		AllScore: []*Score{},
	}

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		// fmt.Println(strings.TrimSpace(e.Text))
		board.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML(".raceItem table", func(e *colly.HTMLElement) {
		count := 0
		score := &Score{}
		e.ForEach("tr th", func(index int, e *colly.HTMLElement) {
			if strings.TrimSpace(e.Text) == "实际" {
				count = index
			}
		})

		e.ForEach(".body-font-color", func(index int, e *colly.HTMLElement) {
			// fmt.Println("Team:", e.Text)
			if index%2 == 0 {
				score.AwayTeam = strings.TrimSpace(e.Text)
			} else {
				score.HomeTeam = strings.TrimSpace(e.Text)
			}
		})

		e.ForEach("tbody tr:nth-child(1)", func(index int, e *colly.HTMLElement) {
			e.ForEach("tbody tr td:nth-child("+fmt.Sprintf("%d", (count+1))+")", func(index int, e *colly.HTMLElement) {
				// fmt.Println("(全场)初盘:", strings.TrimSpace(e.Text))
				score.InitialValue = strings.TrimSpace(e.Text)
			})

			e.ForEach("tbody tr td:nth-child("+fmt.Sprintf("%d", (count+2))+")", func(index int, e *colly.HTMLElement) {
				// fmt.Println("实际:", strings.TrimSpace(e.Text))
				score.FinalValue = strings.TrimSpace(e.Text)
				score.URL = cueerntURL
				board.AllScore = append(board.AllScore, score)
			})
		})
	})

	// On every a element which has href attribute call callback
	c.OnHTML("li a.data-url", func(e *colly.HTMLElement) {
		link := getURL(e.Attr("data-url"))
		e.Request.Visit(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
		cueerntURL = r.URL.String()
		fmt.Println("Visiting", cueerntURL)
	})

	c.OnScraped(func(r *colly.Response) {
	})

	c.Visit(getURL("/bk_league/" + bk + "/p.1?type=ended_race")) // Visit 要放最後

	return board
}
