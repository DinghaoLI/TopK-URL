package utils

import (
	"testing"
)

func Test_MinHeap(t *testing.T) {
	heap := NewMinHeap()
	if heap.Length() == 0 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	heap.Insert(&Url{1, "1"})
	heap.Insert(&Url{2, "2"})
	heap.Insert(&Url{3, "3"})
	heap.Insert(&Url{4, "4"})
	heap.Insert(&Url{5, "5"})

	if heap.Length() == 5 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	v, _ := heap.Min()
	if v.Freq == 1 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	v, _ = heap.DeleteMin()
	if v.Freq == 1 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	v, _ = heap.DeleteMin()
	if v.Freq == 2 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

}
