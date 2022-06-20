package arr

import "math"

// --- binary min heap impl ---
//
type FnComparator func(a, b interface{}) bool

type BinaryHeap struct {
	data  []interface{}
	fcomp FnComparator
}

func NewBinaryHeap(fComp FnComparator) BinaryHeap {
	return BinaryHeap{data: make([]interface{}, 0), fcomp: fComp}
}

func (pq *BinaryHeap) Push(x State) int {
	pq.data = append(pq.data, x)
	pq.siftUp()
	return pq.Size()
}

func (pq *BinaryHeap) Pop() (v interface{}, e bool) {
	if len(pq.data) == 0 {
		return State{}, false
	}
	if len(pq.data) > 1 {
		pq.swap(0, pq.Size()-1)
	}
	v, pq.data = pq.data[pq.Size()-1], pq.data[:pq.Size()-1]
	pq.siftDown()
	return v, true
}

func (pq *BinaryHeap) siftUp() {
	for nodeIdx := pq.Size() - 1; nodeIdx > 0 && pq.comp(nodeIdx, pq.parent(nodeIdx)); {
		pq.swap(nodeIdx, pq.parent(nodeIdx))
		nodeIdx = pq.parent(nodeIdx)
	}
}

func (pq *BinaryHeap) siftDown() {
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

func (pq *BinaryHeap) Size() int     { return len(pq.data) }
func (pq *BinaryHeap) IsEmpty() bool { return pq.Size() == 0 }

func (pq *BinaryHeap) parent(idx int) int { return int(math.Floor(float64(idx-1) / 2)) }
func (pq *BinaryHeap) lidx(idx int) int   { return idx*2 + 1 }
func (pq *BinaryHeap) ridx(idx int) int   { return idx*2 + 2 }
func (pq *BinaryHeap) swap(i, j int)      { pq.data[i], pq.data[j] = pq.data[j], pq.data[i] }
func (pq *BinaryHeap) comp(i, j int) bool { return pq.fcomp(pq.data[i], pq.data[j]) }
