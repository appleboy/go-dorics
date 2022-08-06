package main

import (
	"strconv"
)

type Score struct {
	AwayTeam     string
	HomeTeam     string
	InitialValue string
	FinalValue   string
	URL          string
}

func (s *Score) Outrange() bool {
	initial, _ := strconv.ParseFloat(s.InitialValue, 64)

	if initial >= 28 || initial <= -28 {
		return true
	}

	return false
}

func (s *Score) Outlier() bool {
	initial, _ := strconv.ParseFloat(s.InitialValue, 64)
	final, _ := strconv.ParseFloat(s.FinalValue, 64)
	if initial > 0 && final > 0 {
		return false
	}

	if initial < 0 && final < 0 {
		return false
	}

	if (initial >= 28 && final < 0) ||
		(initial <= -28 && final > 0) {
		return true
	}

	return false
}

type Board struct {
	Title    string
	AllScore []*Score
}
