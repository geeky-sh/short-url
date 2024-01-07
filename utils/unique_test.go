package utils

import (
	"testing"
)

func TestGetShortKey(t *testing.T) {
	var tcases = []struct {
		name string
		size int
	}{
		{"case 1", 5},
		{"case 2", 6},
		{"case 3", 2},
		{"case 4", 4},
		{"case 5", 6},
		{"case 6", 7},
		{"case 7", 8},
		{"case 8", 6},
	}
	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetShortKey(tc.size)
			if len(got) != tc.size {
				t.Errorf("Got size %d Wanted %d\n", len(got), tc.size)
			}
		})
	}
}
