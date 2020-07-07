package easy

import (
	"reflect"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{8, 1, 2, 2, 3}}, []int{4, 0, 1, 1, 3}},
		{"case2", args{[]int{6, 5, 4, 8}}, []int{2, 1, 0, 3}},
		{"case3", args{[]int{7, 7, 7, 7}}, []int{0, 0, 0, 0}},

		{"case4", args{[]int{7}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallerNumbersThanCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
