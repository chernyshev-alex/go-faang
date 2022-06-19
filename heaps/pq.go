package heaps

import (
	"math"
)

type FnComparator func(a, b int) bool
type PriorityQueue struct {
	data  []int
	fcomp FnComparator
}

func NewPriorityQueue() PriorityQueue {
	return NewPriorityQueueWithComp(maxComparator())
}

func NewPriorityQueueWithComp(fc FnComparator) PriorityQueue {
	return PriorityQueue{data: make([]int, 0), fcomp: fc}
}

func (pq *PriorityQueue) Push(x int) int {
	pq.data = append(pq.data, x)
	pq.siftUp()
	return pq.Size()
}

func (pq *PriorityQueue) Peek() (int, bool) {
	if len(pq.data) > 0 {
		var v int
		v, pq.data = pq.data[0], pq.data[1:]
		return v, true
	}
	return 0, false
}

func (pq *PriorityQueue) Pop() (v int, e bool) {
	if len(pq.data) == 0 {
		return 0, false
	}
	if len(pq.data) > 1 {
		pq.swap(0, pq.Size()-1)
	}
	v, pq.data = pq.data[pq.Size()-1], pq.data[:pq.Size()-1]
	pq.siftDown()
	return v, true
}

func (pq *PriorityQueue) siftUp() {
	for nodeIdx := pq.Size() - 1; nodeIdx > 0 && pq.comp(nodeIdx, pq.parent(nodeIdx)); {
		pq.swap(nodeIdx, pq.parent(nodeIdx))
		nodeIdx = pq.parent(nodeIdx)
	}
}

func (pq *PriorityQueue) siftDown() {
	for nodeIdx, sz := 0, len(pq.data); (pq.lidx(nodeIdx) < sz && pq.comp(pq.lidx(nodeIdx), nodeIdx)) ||
		(pq.ridx(nodeIdx) < sz && pq.comp(pq.ridx(nodeIdx), nodeIdx)); {

		greaterChildIdx := pq.lidx(nodeIdx)
		if pq.ridx(nodeIdx) < sz && pq.comp(pq.ridx(nodeIdx), pq.lidx(nodeIdx)) {
			greaterChildIdx = pq.ridx(nodeIdx)
		}
		pq.swap(greaterChildIdx, nodeIdx)
		nodeIdx = greaterChildIdx

	}
}

func (pq *PriorityQueue) Size() int     { return len(pq.data) }
func (pq *PriorityQueue) IsEmpty() bool { return pq.Size() == 0 }

func (pq *PriorityQueue) parent(idx int) int { return int(math.Floor(float64(idx-1) / 2)) }
func (pq *PriorityQueue) lidx(idx int) int   { return idx*2 + 1 }
func (pq *PriorityQueue) ridx(idx int) int   { return idx*2 + 2 }
func (pq *PriorityQueue) swap(i, j int)      { pq.data[i], pq.data[j] = pq.data[j], pq.data[i] }
func (pq *PriorityQueue) comp(i, j int) bool { return pq.fcomp(pq.data[i], pq.data[j]) }

func maxComparator() FnComparator { return func(a, b int) bool { return a > b } }
