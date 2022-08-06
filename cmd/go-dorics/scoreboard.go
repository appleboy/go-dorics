package main

import "strconv"

type Score struct {
	AwayTeam     string
	HomeTeam     string
	InitialValue string
	FinalValue   string
	URL          string
	IsOutlier    bool
}

func (s *Score) Outlier() {
	initial, _ := strconv.ParseFloat(s.InitialValue, 64)
	final, _ := strconv.ParseFloat(s.FinalValue, 64)
	if initial > 0 && final > 0 {
		return
	}

	if initial < 0 && final < 0 {
		return
	}

	if (initial >= 28 && final < 0) ||
		(initial <= 28 && final > 0) {
		s.IsOutlier = true
	}

	return
}

type Board struct {
	Title    string
	AllScore []*Score
}
