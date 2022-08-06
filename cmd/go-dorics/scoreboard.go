package main

type Score struct {
	AwayTeam     string
	HomeTeam     string
	InitialValue string
	FinalValue   string
	URL          string
}

type Board struct {
	Title    string
	AllScore []*Score
}
