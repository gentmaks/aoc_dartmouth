// Package models
package models

type FloatHeap []float64

func (h FloatHeap) Len() int { return len(h) }

func (h FloatHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h FloatHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *FloatHeap) Push(x any) {
	*h = append(*h, x.(float64))
}

func (h *FloatHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
