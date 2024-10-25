package lgo

import (
	"container/heap"
	"math"
	"sort"
	"strings"
)

// 1942. The Number of the Smallest Unoccupied Chair - med

func smallestChair1(times [][]int, targetFriend int) int {
	var arrival timelines
	var leaving timelines
	for idx, t := range times {
		arrival = append(arrival, timeline{t[0], idx})
		leaving = append(leaving, timeline{t[1], idx})
	}

	sort.Sort(arrival)
	sort.Sort(leaving)
	// fmt.Println(arrival)
	// fmt.Println(leaving)

	var sitting int
	var seatBacklog []int
	var seatass = make(map[int]int)
	i, j := 0, 0
	for {
		a := arrival[i]
		l := leaving[j]
		// fmt.Println(a, l)
		if a.time < l.time {
			// fmt.Println("arrival")
			// arrival

			// find smallest seat first
			var smallestSeat int
			if len(seatBacklog) == 0 {
				smallestSeat = sitting
			} else {
				smallestSeat = seatBacklog[0]
				seatBacklog = seatBacklog[1:]
			}

			if a.friendIdx == targetFriend {
				return smallestSeat
			}
			seatass[a.friendIdx] = smallestSeat

			i++
			sitting++
		} else {
			// fmt.Println("leaving")
			// leaving
			seatIdx := seatass[l.friendIdx]
			seatBacklog = append(seatBacklog, seatIdx)
			sort.Slice(seatBacklog, func(i, j int) bool {
				return seatBacklog[i] < seatBacklog[j]
			})

			j++
			sitting--
		}
	}

	return 0
}

type timelines []timeline

type timeline struct {
	time      int
	friendIdx int
}

func (a timelines) Less(i, j int) bool {
	return a[i].time < a[j].time
}

func (a timelines) Len() int {
	return len(a)
}

func (a timelines) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 2530. Maximal Score After Applying K Operations

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return x
}

func maxKelements(nums []int, k int) int64 {
	h := IntHeap(nums)

	heap.Init(&h)

	result := int64(0)
	for k > 0 {
		x := heap.Pop(&h).(int)
		result += int64(x)
		heap.Push(&h, int(math.Ceil(float64(x)/3.0)))
		k--
	}

	return result
}

// 1405. Longest Happy String
type CharCount struct {
	count int
	ch    byte
}

// Priority queue implementation for Go
type MaxHeap []CharCount

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(CharCount))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func longestDiverseString(a int, b int, c int) string {
	pq := &MaxHeap{}
	heap.Init(pq)

	if a > 0 {
		heap.Push(pq, CharCount{a, 'a'})
	}
	if b > 0 {
		heap.Push(pq, CharCount{b, 'b'})
	}
	if c > 0 {
		heap.Push(pq, CharCount{c, 'c'})
	}

	var result strings.Builder

	for pq.Len() > 0 {
		first := heap.Pop(pq).(CharCount)

		if result.Len() >= 2 && result.String()[result.Len()-1] == first.ch && result.String()[result.Len()-2] == first.ch {
			if pq.Len() == 0 {
				break
			}
			second := heap.Pop(pq).(CharCount)
			result.WriteByte(second.ch)
			second.count--
			if second.count > 0 {
				heap.Push(pq, second)
			}
			heap.Push(pq, first)
		} else {
			result.WriteByte(first.ch)
			first.count--
			if first.count > 0 {
				heap.Push(pq, first)
			}
		}
	}

	return result.String()
}
