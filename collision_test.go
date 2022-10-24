package main

import "testing"

func Test_collisionDetection(t *testing.T) {

	type args struct {
		x1 int
		y1 int
		w1 int
		h1 int

		x2 int
		y2 int
		w2 int
		h2 int

		t int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "0",
			args: args{0, 0, 1, 1, 0, 0, 1, 1, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collisionDetection(
				tt.args.x1,
				tt.args.y1,
				tt.args.w1,
				tt.args.h1,

				tt.args.x2,
				tt.args.y2,
				tt.args.w2,
				tt.args.h2,

				tt.args.t,
			); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
