package main

import "testing"

func TestNormalization(t *testing.T) {
	testCases := []struct {
		Input  string
		Output string
	}{
		{"1234567890", "(123) 4567890"},
		{"123 456 7891", "(123) 4567891"},
		{"(123) 456 7892", "(123) 4567892"},
		{"(123) 456-7893", "(123) 4567893"},
		{"123-456-7894", "(123) 4567894"},
		{"123-456-7890", "(123) 4567890"},
		{"(123)456-7892", "(123) 4567892"},
	}
	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			want := Normalization(tc.Input)
			if want != tc.Output {
				t.Errorf("got %s; want %s", tc.Output, want)
			}
		})
	}

}
