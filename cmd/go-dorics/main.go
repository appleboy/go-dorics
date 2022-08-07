package main

import (
	"fmt"
	"strings"

	"github.com/appleboy/go-dorics/pkg/analytic"

	"github.com/xuri/excelize/v2"
)

func main() {
	board := analytic.Spider("111")
	scoreTyes := []interface{}{}
	scoreCounts := []interface{}{}
	scoreExceptCounts := []interface{}{}
	scoreExceptDetails := []interface{}{}
	for _, v := range board.Export() {
		fmt.Println()
		fmt.Println("score type:", v.Score, "Outrange Count:", len(v.OutrangeList), "Outlier Count:", len(v.OutlierList))
		scoreTyes = append(scoreTyes, v.Score)
		scoreCounts = append(scoreCounts, len(v.OutrangeList))
		scoreExceptCounts = append(scoreExceptCounts, len(v.OutlierList))
		fmt.Println("======Outrange Start=========")
		excepts := []string{}
		for _, o := range v.OutrangeList {
			fmt.Printf("%s vs %s\n", o.AwayTeam, o.HomeTeam)
		}
		fmt.Println("======Outrange End=========")
		if len(v.OutlierList) > 0 {
			fmt.Println("======Outlier Start=========")
			for _, o := range v.OutlierList {
				fmt.Printf("%s vs %s\n", o.AwayTeam, o.HomeTeam)
				excepts = append(excepts, fmt.Sprintf("%s vs %s (%s)", o.AwayTeam, o.HomeTeam, o.URL))
				excepts = append(excepts, fmt.Sprintf("%s vs %s (%s)", o.AwayTeam, o.HomeTeam, o.URL))
			}
			fmt.Println("======Outlier End=========")
		}
		scoreExceptDetails = append(scoreExceptDetails, strings.Join(excepts, " and "))
	}

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	f.SetCellValue("Sheet1", "B1", board.Title)
	f.SetCellValue("Sheet1", "B3", "全部幾場")
	f.SetCellValue("Sheet1", "B4", "例外")
	f.SetCellValue("Sheet1", "B5", "例外明細")

	f.SetSheetRow("Sheet1", "C2", &scoreTyes)
	f.SetSheetRow("Sheet1", "C3", &scoreCounts)
	f.SetSheetRow("Sheet1", "C4", &scoreExceptCounts)
	f.SetSheetRow("Sheet1", "C5", &scoreExceptDetails)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
