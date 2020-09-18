package grotain_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/tiagoemsi/grotain"
)

func TestCalcInterval(test *testing.T) {
	scenarios := []struct {
		name string
		item grotain.Cards
		want int
	}{
		{"CalcInterval - Stage 1", grotain.Cards{Interval: 0, Efactor: 2.5}, 1},
		{"CalcInterval - Stage 2", grotain.Cards{Interval: 1, Efactor: 2.5}, 6},
		{"CalcInterval - Stage 3", grotain.Cards{Interval: 6, Efactor: 2.5}, 13},
		{"CalcInterval - Stage 4", grotain.Cards{Interval: 13, Efactor: 2.5}, 30},
		{"CalcInterval - Stage 5", grotain.Cards{Interval: 30, Efactor: 2.5}, 73},
		{"CalcInterval - Stage 6", grotain.Cards{Interval: 73, Efactor: 2.5}, 180},
		{"CalcInterval - Stage 7", grotain.Cards{Interval: 180, Efactor: 2.5}, 448},
	}

	for i, s := range scenarios {
		test.Run(
			fmt.Sprintf("[%d]-%s", i, s.name),
			func(t *testing.T) {
				got := s.item.CalcInterval(4)
				if got != s.want {
					t.Errorf("\ngot: %v\nwant: %v", got, s.want)
				}
			},
		)
	}
}

func TestCalcEfactor(test *testing.T) {
	scenarios := []struct {
		name    string
		item    grotain.Cards
		quality int
		want    float64
	}{
		{"CalcEfactor - Stage 1", grotain.Cards{Efactor: 2.5}, 0, 1.7000000000000002},
		{"CalcEfactor - Stage 2", grotain.Cards{Efactor: 2.5}, 1, 1.96},
		{"CalcEfactor - Stage 3", grotain.Cards{Efactor: 2.5}, 2, 2.1799999999999997},
		{"CalcEfactor - Stage 4", grotain.Cards{Efactor: 2.5}, 3, 2.36},
		{"CalcEfactor - Stage 5", grotain.Cards{Efactor: 2.5}, 4, 2.5},
		{"CalcEfactor - Stage 6", grotain.Cards{Efactor: 2.5}, 5, 2.6},
	}

	for i, s := range scenarios {
		test.Run(
			fmt.Sprintf("[%d]-%s", i, s.name),
			func(t *testing.T) {
				got := s.item.CalcEfactor(s.quality)
				if got != s.want {
					t.Errorf("\ngot: %v\nwant: %v", got, s.want)
				}
			},
		)
	}
}

func TestReviewEfactor(test *testing.T) {
	scenarios := []struct {
		name    string
		item    grotain.Cards
		quality int
		want    grotain.Cards
	}{
		{"Review - Quality 0", grotain.Cards{Interval: 6980, Efactor: 2}, 0, grotain.Cards{Interval: 1, Efactor: 1.3}},
		{"Review - Quality 1", grotain.Cards{Interval: 6980, Efactor: 2}, 1, grotain.Cards{Interval: 1, Efactor: 1.46}},
		{"Review - Quality 2", grotain.Cards{Interval: 6980, Efactor: 2}, 2, grotain.Cards{Interval: 1, Efactor: 1.68}},
		{"Review - Quality 3", grotain.Cards{Interval: 6980, Efactor: 2}, 3, grotain.Cards{Interval: 12981, Efactor: 1.86}},
		{"Review - Quality 4", grotain.Cards{Interval: 6980, Efactor: 2}, 4, grotain.Cards{Interval: 13958, Efactor: 2}},
		{"Review - Quality 5", grotain.Cards{Interval: 6980, Efactor: 2}, 5, grotain.Cards{Interval: 14656, Efactor: 2.1}},
	}

	for i, s := range scenarios {
		test.Run(
			fmt.Sprintf("[%d]-%s", i, s.name),
			func(t *testing.T) {
				got := s.item.Review(s.quality)
				if !reflect.DeepEqual(got, &s.want) {
					t.Errorf("\ngot: %v\nwant: %v", got, &s.want)
				}
			},
		)
	}
}
func TestReviewInterval(test *testing.T) {
	scenarios := []struct {
		name    string
		item    grotain.Cards
		quality int
		want    grotain.Cards
	}{
		{"Review - 10x Quality 0", grotain.Cards{Interval: 0, Efactor: 2.5}, 0, grotain.Cards{Interval: 1, Efactor: 1.3}},
		{"Review - 10x Quality 1", grotain.Cards{Interval: 0, Efactor: 2.5}, 1, grotain.Cards{Interval: 1, Efactor: 1.3}},
		{"Review - 10x Quality 2", grotain.Cards{Interval: 0, Efactor: 2.5}, 2, grotain.Cards{Interval: 1, Efactor: 1.3}},
		{"Review - 10x Quality 3", grotain.Cards{Interval: 0, Efactor: 2.5}, 3, grotain.Cards{Interval: 196, Efactor: 1.3}},
		{"Review - 10x Quality 4", grotain.Cards{Interval: 0, Efactor: 2.5}, 4, grotain.Cards{Interval: 17448, Efactor: 2.5}},
		{"Review - 10x Quality 5", grotain.Cards{Interval: 0, Efactor: 2.5}, 5, grotain.Cards{Interval: 153605, Efactor: 3.600000000000001}},
	}
	for i, s := range scenarios {
		test.Run(
			fmt.Sprintf("[%d]-%s", i, s.name),
			func(t *testing.T) {
				var got *grotain.Cards
				for j := 0; j <= 10; j++ {
					got = s.item.Review(s.quality)
				}
				if !reflect.DeepEqual(got, &s.want) {
					t.Errorf("\ngot: %v\nwant: %v", got, &s.want)
				}
			},
		)
	}
}
