package analytic

import (
	"sort"
	"strconv"
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

func (b *Board) Excel() []*Excel {
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
