package lcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateList(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "equal",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
		{
			name: "empty",
			args: args{
				arr: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, generateList(tt.args.arr), "generateList(%v)", tt.args.arr)
		})
	}
}
