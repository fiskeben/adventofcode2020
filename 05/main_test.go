package main

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		expr  string
		s     space
		front rune
		back  rune
	}
	tests := []struct {
		name string
		arg  args
		want int
	}{
		{
			name: "finds row number",
			arg: args{
				expr:  "FBFBBFF",
				s:     space{0, 127},
				front: 'F',
				back:  'B',
			},
			want: 44,
		},
		{
			name: "finds col number",
			arg: args{
				expr:  "RLR",
				s:     space{0, 7},
				front: 'L',
				back:  'R',
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := divide(tt.arg.expr, tt.arg.s, tt.arg.front, tt.arg.back)
			if res != tt.want {
				t.Errorf("expected %d got %d", tt.want, res)
			}
		})
	}
}
