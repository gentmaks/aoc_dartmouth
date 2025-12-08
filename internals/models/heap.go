// Package models
package models

type Vector struct {
	Dist   float64
	Source int
	Target int
}

type VectorHeap []Vector

func (h VectorHeap) Len() int { return len(h) }

func (h VectorHeap) Less(i, j int) bool { return h[i].Dist < h[j].Dist }

func (h VectorHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *VectorHeap) Push(x any) {
	*h = append(*h, x.(Vector))
}

func (h *VectorHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
