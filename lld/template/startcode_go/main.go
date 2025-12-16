package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Top() any {
	old := *h
	n := len(old)
	x := old[n-1]
	return x
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Top() any {
	old := *h
	n := len(old)
	x := old[n-1]
	return x
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	MnHeap *MinHeap
	MxHeap *MaxHeap
}

func Constructor() MedianFinder {
	mih := &MinHeap{}
	mxh := &MaxHeap{}
	heap.Init(mih)
	heap.Init(mxh)
	return MedianFinder{
		MnHeap: mih,
		MxHeap: mxh,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.MnHeap.Len() == this.MxHeap.Len() {
		heap.Push(this.MxHeap, num)
		this.MnHeap.Push(heap.Pop(this.MxHeap))
	} else {
		heap.Push(this.MxHeap, heap.Pop(this.MnHeap))
		heap.Push(this.MxHeap, num)
		heap.Push(this.MnHeap, heap.Pop(this.MxHeap))
	}
}
func (this *MedianFinder) FindMedian() float64 {
	left := 0
	right := 0
	if this.MxHeap.Len() > 0 {
		left = this.MxHeap.Top().(int)
	}

	if this.MnHeap.Len() > 0 {
		right = this.MnHeap.Top().(int)
	}

	if this.MnHeap.Len() == this.MxHeap.Len() {
		return float64(right+left) / 2.0
	}

	return float64(right)

}
func main() {

}
