package main

import (
	"fmt"

	"github.com/appleboy/go-dorics/pkg/analytic"
)

func main() {
	board := analytic.Spider("383")
	for _, v := range board.Export() {
		fmt.Println()
		fmt.Println("score type:", v.Score, "Outrange Count:", len(v.OutrangeList), "Outlier Count:", len(v.OutlierList))
		fmt.Println("======Outrange Start=========")
		for _, o := range v.OutrangeList {
			fmt.Printf("%s vs %s\n", o.AwayTeam, o.HomeTeam)
		}
		fmt.Println("======Outrange End=========")
		if len(v.OutlierList) > 0 {
			fmt.Println("======Outlier Start=========")
			for _, o := range v.OutlierList {
				fmt.Printf("%s vs %s\n", o.AwayTeam, o.HomeTeam)
			}
			fmt.Println("======Outlier End=========")
		}
	}
}
