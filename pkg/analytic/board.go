package analytic

import "strconv"

type Excel struct {
	OutlierList  []*Score
	OutrangeList []*Score
}

type Board struct {
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
			b.Data[initial] = &Excel{}
		}
		b.Data[initial].OutrangeList = append(b.Data[initial].OutrangeList, v)
		if v.Outlier() {
			b.Data[initial].OutlierList = append(b.Data[initial].OutlierList, v)
		}
	}
}
