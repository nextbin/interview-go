package easy

import "testing"

func Test_xorOperation(t *testing.T) {
	type args struct {
		n     int
		start int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{5, 0}, 8},
		{"case2", args{4, 3}, 8},
		{"case3", args{1, 7}, 7},
		{"case4", args{10, 5}, 2},

		{"case5", args{1, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorOperation(tt.args.n, tt.args.start); got != tt.want {
				t.Errorf("xorOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
