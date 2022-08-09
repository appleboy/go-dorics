package analytic

import (
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Excel struct {
	Score        float64
	OutlierList  []*Score
	OutrangeList []*Score
}

type Board struct {
	ID       string
	Title    string
	AllScore []*Score
	Data     map[float64]*Excel
}

func (b *Board) Output() {
	b.Data = make(map[float64]*Excel)
	for _, v := range b.AllScore {
		if !v.Outrange() {
			continue
		}
		initial, _ := strconv.ParseFloat(v.InitialValue, 64)
		if _, ok := b.Data[initial]; !ok {
			b.Data[initial] = &Excel{
				Score: initial,
			}
		}
		b.Data[initial].OutrangeList = append(b.Data[initial].OutrangeList, v)
		if v.Outlier() {
			b.Data[initial].OutlierList = append(b.Data[initial].OutlierList, v)
		}
	}
}

func (b *Board) Export() []*Excel {
	if b.Data == nil {
		b.Output()
	}

	if len(b.Data) == 0 {
		return nil
	}
	keys := make([]float64, 0, len(b.Data))
	for k := range b.Data {
		keys = append(keys, k)
	}

	sort.Float64s(keys)
	data := make([]*Excel, 0, len(b.Data))
	for _, k := range keys {
		data = append(data, b.Data[k])
	}

	return data
}

func (b *Board) Excel(folder string) error {
	scoreTyes := []interface{}{}
	scoreCounts := []interface{}{}
	scoreExceptCounts := []interface{}{}
	scoreExceptDetails := []interface{}{}
	for _, v := range b.Export() {
		scoreTyes = append(scoreTyes, v.Score)
		scoreCounts = append(scoreCounts, len(v.OutrangeList))
		scoreExceptCounts = append(scoreExceptCounts, len(v.OutlierList))
		excepts := []string{}
		if len(v.OutlierList) > 0 {
			for _, o := range v.OutlierList {
				excepts = append(excepts, fmt.Sprintf("%s vs %s (%s)", o.AwayTeam, o.HomeTeam, o.URL))
			}
		}
		scoreExceptDetails = append(scoreExceptDetails, strings.Join(excepts, " and "))
	}

	output := path.Join(folder, b.ID+".xlsx")

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	f.SetCellValue("Sheet1", "B1", b.Title)
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
	return f.SaveAs(output)
}
