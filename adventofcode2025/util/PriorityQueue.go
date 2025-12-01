package util

import (
	"container/heap"
	"slices"
)

type Item interface {
	Index() int
	SetIndex(index int)

	Update(i Item)

	Less(i Item) bool
}

type PriorityQueue[T Item] []T

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Less(pq[j])
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].SetIndex(i)
	pq[j].SetIndex(j)
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)

	item := x.(T)

	item.SetIndex(n)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]

	slices.Delete(old, n-1, n-1)
	item.SetIndex(-1)
	*pq = old[0 : n-1]
	return item
}
func (pq *PriorityQueue[T]) Update(oldItem T, newItem T) {
	oldItem.Update(newItem)
	heap.Fix(pq, oldItem.Index())
}
