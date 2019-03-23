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

// MinHeap构造方法
func NewMinHeap() *MinHeap {
    // 第一个元素仅用于结束insert中的 for 循环
    first := &Url{math.MinInt64, "None"}
    h := &MinHeap{Element: []*Url{first}}
    return h
}

// 堆的大小
func (H *MinHeap) Length() int {
	return len(H.Element) - 1
}

// 获取最小堆的最小值
func (H *MinHeap) Min() (*Url, error) {
	if len(H.Element) > 1 {
		return H.Element[1], nil
	}
	return nil, fmt.Errorf("heap is empty")
}

// MinHeap格式化输出
// func (H *MinHeap) String() string {
// 	return fmt.Sprint64f("%v", H.Element[1:])
// }

// 插入数字,插入数字需要保证堆的性质
func (H *MinHeap) Insert(v *Url) {
    H.Element = append(H.Element, v)
    i := len(H.Element) - 1
    // 上浮
    for ; (H.Element[i/2]).Freq > v.Freq; i /= 2 {
        H.Element[i] = H.Element[i/2]
    }

    H.Element[i] = v
}

// 删除并返回最小值
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
			child ++
		}
		// 下滤一层
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