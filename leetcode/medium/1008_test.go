package medium

import (
	"testing"
	"fmt"
)

func Test_bstFromPreorder(t *testing.T) {
	type args struct {
		preorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"case1",args{[]int{8,5,1,7,10,12}},nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bstFromPreorder(tt.args.preorder)
			fmt.Println(got)
			fmt.Println(got.Val)
			fmt.Println(got.Left.Val)
			fmt.Println(got.Right.Val)
		})
	}
}

