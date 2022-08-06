package main

import "testing"

func TestScore_Outlier(t *testing.T) {
	type fields struct {
		AwayTeam     string
		HomeTeam     string
		InitialValue string
		FinalValue   string
		URL          string
		IsOutlier    bool
	}
	tests := []struct {
		name    string
		fields  fields
		outlier bool
	}{
		{
			name: "outlier",
			fields: fields{
				InitialValue: "+28.5",
				FinalValue:   "-10.0",
			},
			outlier: true,
		},
		{
			name: "not outlier",
			fields: fields{
				InitialValue: "+20",
				FinalValue:   "+10",
			},
			outlier: false,
		},
		{
			name: "not outlier",
			fields: fields{
				InitialValue: "+20",
				FinalValue:   "-10",
			},
			outlier: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Score{
				AwayTeam:     tt.fields.AwayTeam,
				HomeTeam:     tt.fields.HomeTeam,
				InitialValue: tt.fields.InitialValue,
				FinalValue:   tt.fields.FinalValue,
				URL:          tt.fields.URL,
				IsOutlier:    tt.fields.IsOutlier,
			}
			s.Outlier()
			if s.IsOutlier != tt.outlier {
				t.Fatalf("want: %v, current outlier is %v", tt.outlier, s.IsOutlier)
			}
		})
	}
}
