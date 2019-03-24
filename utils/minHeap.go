package utils

import (
	"fmt"
	"math"
)

type MinHeap struct {
	Element []*Url
}

type Url struct {
	Freq int64
	Addr string
}

// MinHeap constructor
func NewMinHeap() *MinHeap {
	first := &Url{math.MinInt64, "None"}
	h := &MinHeap{Element: []*Url{first}}
	return h
}

// Length of Minheap
func (H *MinHeap) Length() int {
	return len(H.Element) - 1
}

// Get the minimum of the Minheap
func (H *MinHeap) Min() (*Url, error) {
	if len(H.Element) > 1 {
		return H.Element[1], nil
	}
	return nil, fmt.Errorf("heap is empty")
}

// Inserting items requires ensuring the nature of the Minheap
func (H *MinHeap) Insert(v *Url) {
	H.Element = append(H.Element, v)
	i := len(H.Element) - 1
	for ; (H.Element[i/2]).Freq > v.Freq; i /= 2 {
		H.Element[i] = H.Element[i/2]
	}

	H.Element[i] = v
}

// Delete and return the minimum
func (H *MinHeap) DeleteMin() (*Url, error) {
	if len(H.Element) <= 1 {
		return nil, fmt.Errorf("MinHeap is empty")
	}
	minElement := H.Element[1]
	lastElement := H.Element[len(H.Element)-1]
	var i, child int
	for i = 1; i*2 < len(H.Element); i = child {
		child = i * 2
		if child < len(H.Element)-1 && H.Element[child+1].Freq < H.Element[child].Freq {
			child++
		}
		if lastElement.Freq > H.Element[child].Freq {
			H.Element[i] = H.Element[child]
		} else {
			break
		}
	}
	H.Element[i] = lastElement
	H.Element = H.Element[:len(H.Element)-1]
	return minElement, nil
}
