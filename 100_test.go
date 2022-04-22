package lcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test100(t *testing.T) {
	isSameTree := func(p *TreeNode, q *TreeNode) bool {
		return false
	}
	isSameTree = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	assert.Equal(t, isSameTree(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}), true)

	assert.Equal(t, isSameTree(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		}), false)

	assert.Equal(t, isSameTree(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		}), false)

}
