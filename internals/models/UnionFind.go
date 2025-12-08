// Package models
package models

type UnionFind struct {
	parents []int
	size    []int
}

func UFInit(n int) *UnionFind {
	parents := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		size[i] = 1
	}
	return &UnionFind{parents: parents, size: size}
}

func (uf *UnionFind) Find(node int) int {
	if uf.parents[node] != node {
		uf.parents[node] = uf.Find(uf.parents[node])
	}
	return uf.parents[node]
}

func (uf *UnionFind) Union(node1 int, node2 int) bool {
	node1P := uf.Find(node1)
	node2P := uf.Find(node2)
	if node1P == node2P {
		return false
	}
	if uf.size[node1P] > uf.size[node2P] {
		uf.parents[node2P] = node1P
	} else if uf.size[node1P] < uf.size[node2P] {
		uf.parents[node1P] = node2P
	} else {
		uf.parents[node1P] = node2P
		uf.size[node2P] += uf.size[node1P]
	}
	return true
}

func (uf *UnionFind) GetConnCompCount() int {
	var count int
	for i, p := range uf.parents {
		if i == p {
			count++
		}
	}
	return count
}
