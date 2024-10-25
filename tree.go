package lgo

import "container/heap"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 94. Binary Tree Inorder Traversal

func inorderTraversal(root *TreeNode) []int {
	var result []int

	var iT func(root *TreeNode)
	iT = func(root *TreeNode) {
		if root == nil {
			return
		}
		iT(root.Left)
		result = append(result, root.Val)
		iT(root.Right)
	}
	iT(root)
	return result
}

// 100. Same Tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 101. Symmetric Tree

func isSymmetric(root *TreeNode) bool {
	return isSameSymmetricTree(root.Left, root.Right)
}

func isSameSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
}

// 2583. Kth Largest Sum in a Binary Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type qu []int

func (h qu) Len() int           { return len(h) }
func (h qu) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h qu) Less(i, j int) bool { return h[i] < h[j] }

func (h *qu) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *qu) Pop() any {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := []*TreeNode{root}
	minHeap := new(qu)
	heap.Init(minHeap)
	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		for i := 0; i < n; i++ {
			node := queue[i]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if minHeap.Len() == k && (*minHeap)[0] < sum {
			heap.Pop(minHeap)
			heap.Push(minHeap, sum)
		}
		if minHeap.Len() < k {
			heap.Push(minHeap, sum)
		}
		queue = queue[n:]
	}
	if minHeap.Len() != k {
		return -1
	}
	return int64((*minHeap)[0])
}
