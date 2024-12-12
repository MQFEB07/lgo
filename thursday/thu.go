package thursday

import (
	"container/heap"
	"math"
)

// 2558. Take Gifts From the Richest Pile

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	h := MaxHeap(gifts)

	heap.Init(&h)

	for _ = range k {
		x := heap.Pop(&h).(int)
		x = int(math.Sqrt(float64(x)))
		heap.Push(&h, x)
	}
	remain := 0
	for h.Len() > 0 {
		remain += heap.Pop(&h).(int)
	}
	return int64(remain)
}
