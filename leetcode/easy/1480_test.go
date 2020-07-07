package easy

import (
	"reflect"
	"testing"
)

func Test_runningSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 2, 3, 4}}, []int{1, 3, 6, 10}},
		{"case2", args{[]int{1, 1, 1, 1, 1}}, []int{1, 2, 3, 4, 5}},
		{"case3", args{[]int{3, 1, 2, 10, 1}}, []int{3, 4, 6, 16, 17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
