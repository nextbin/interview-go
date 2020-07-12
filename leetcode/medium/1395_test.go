package medium

import "testing"

func Test_numTeams(t *testing.T) {
	type args struct {
		rating []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{2, 5, 3, 4, 1}}, 3},
		{"case2", args{[]int{2, 1, 3}}, 0},
		{"case3", args{[]int{1, 2, 3, 4}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTeams(tt.args.rating); got != tt.want {
				t.Errorf("numTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
