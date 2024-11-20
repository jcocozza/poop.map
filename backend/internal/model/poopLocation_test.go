package model

import (
	"fmt"
	"testing"
)

var maskTests = []struct {
	mask    int
	seasons []Season
}{
	{0, []Season{}},

	{1, []Season{Summer}},
	{2, []Season{Fall}},
	{4, []Season{Winter}},
	{8, []Season{Spring}},

	{1 + 2, []Season{Summer, Fall}},
	{1 + 4, []Season{Summer, Winter}},
	{1 + 8, []Season{Summer, Spring}},

	{2 + 4, []Season{Fall, Winter}},
	{2 + 8, []Season{Fall, Spring}},

	{4 + 8, []Season{Winter, Spring}},

	{1 + 2 + 4, []Season{Summer, Fall, Winter}},
	{1 + 2 + 8, []Season{Summer, Fall, Spring}},
	{1 + 4 + 8, []Season{Summer, Winter, Spring}},
	{2 + 4 + 8, []Season{Fall, Winter, Spring}},

	{1 + 2 + 4 + 8, []Season{Summer, Fall, Winter, Spring}},
}

func TestGetSeasons(t *testing.T) {
	for _, tt := range maskTests {
		t.Run(fmt.Sprint(tt.mask), func(t *testing.T) {
			seasons := GetSeasons(tt.mask)
			if len(seasons) != len(tt.seasons) {
				t.Errorf("got %v, want %v", seasons, tt.seasons)
			}
			for i := range seasons {
				if seasons[i] != tt.seasons[i] {
					t.Errorf("got %s, want %s", seasons[i], tt.seasons[i])
				}
			}
		})
	}
}

func TestSeasonMask(t *testing.T) {
	for _, tt := range maskTests {
		t.Run(fmt.Sprint(tt.mask), func(t *testing.T) {
			mask := SeasonMask(tt.seasons)
			if mask != tt.mask {
				t.Errorf("got %d, want %d", mask, tt.mask)
			}
		})
	}
}
