package analytic

import (
	"testing"
)

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
		name   string
		fields fields
		want   bool
	}{
		{
			name: "outlier",
			fields: fields{
				InitialValue: "+28.5",
				FinalValue:   "-10.0",
			},
			want: true,
		},
		{
			name: "not outlier",
			fields: fields{
				InitialValue: "+20",
				FinalValue:   "+10",
			},
			want: false,
		},
		{
			name: "not outlier",
			fields: fields{
				InitialValue: "+20",
				FinalValue:   "-10",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Score{
				InitialValue: tt.fields.InitialValue,
				FinalValue:   tt.fields.FinalValue,
			}
			if got := s.Outlier(); got != tt.want {
				t.Errorf("Score.Outlier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScore_Outrange(t *testing.T) {
	type fields struct {
		AwayTeam     string
		HomeTeam     string
		InitialValue string
		FinalValue   string
		URL          string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "outrange",
			fields: fields{
				InitialValue: "+28.5",
				FinalValue:   "-10.0",
			},
			want: true,
		},
		{
			name: "not outrange",
			fields: fields{
				InitialValue: "+20",
				FinalValue:   "+10",
			},
			want: false,
		},
		{
			name: "not outrange",
			fields: fields{
				InitialValue: "+20",
				FinalValue:   "-10",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Score{
				InitialValue: tt.fields.InitialValue,
				FinalValue:   tt.fields.FinalValue,
			}
			if got := s.Outrange(); got != tt.want {
				t.Errorf("Score.Outrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
