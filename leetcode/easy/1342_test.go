package easy

import "testing"

func Test_numberOfSteps(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{14}, 6},
		{"case2", args{8}, 4},
		{"case3", args{123}, 12},

		{"case4", args{0}, 0},
		{"case5", args{1}, 1},
		{"case6", args{2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSteps(tt.args.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
