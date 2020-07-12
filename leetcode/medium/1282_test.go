package medium

import (
	"reflect"
	"testing"
)

func Test_groupThePeople(t *testing.T) {
	type args struct {
		groupSizes []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{[]int{3, 3, 3, 3, 3, 1, 3}}, [][]int{{0, 1, 2}, {3, 4, 6}, {5}}},
		{"case2", args{[]int{2, 1, 3, 3, 3, 2}}, [][]int{{0, 5}, {1}, {2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupThePeople(tt.args.groupSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupThePeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
