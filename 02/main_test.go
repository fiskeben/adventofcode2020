package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_parsePassword(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		want    password
	}{
		{
			name:  "parses password",
			input: "1-3 a: abcde",
			want: password{
				min:  1,
				max:  3,
				char: 'a',
				pass: "abcde",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := parsePassword(tt.input)
			if tt.wantErr != (err != nil) {
				t.Fatalf("wanted err=%t", tt.wantErr)
			}

			if diff := cmp.Diff(tt.want, res); diff != "" {
				t.Error(diff)
			}
		})
	}
}
