package arrays

import (
	"container/heap"
	"fmt"
)

type freqElement struct {
	value int
	count int
}

type pq []freqElement

func (h pq) Len() int {
	return len(h)
}

func (h pq) Less(i, j int) bool {
	// for max heap
	return h[i].count > h[j].count
}

func (h pq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pq) Push(x any) {
	*h = append(*h, x.(freqElement))
}

func (h *pq) Pop() any {
	old, n := *h, len(*h)
	popped := old[n-1]
	*h = old[:n-1]
	return popped
}

func TopKFreqElements() {

	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	ans := topKFreqElements(nums, k)
	fmt.Println(ans)

}

func topKFreqElements(nums []int, k int) []int {

	if k >= len(nums) {
		return nums
	}

	numsMap := make(map[int]int)

	for _, num := range nums {
		numsMap[num] += 1
	}

	maxHeap := &pq{}

	for val, count := range numsMap {
		if maxHeap.Len() < k {
			maxHeap.Push(freqElement{val, count})
			if maxHeap.Len() == k {
				heap.Init(maxHeap)
			}
		} else {
			heap.Push(maxHeap, freqElement{val, count})
		}
	}

	var res []int
	for i := 0; i < k; i++ {

		eVal := heap.Pop(maxHeap).(freqElement)
		res = append(res, eVal.value)
	}

	return res

}
