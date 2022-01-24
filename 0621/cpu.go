package leetcode

/*
这道题， 看官方题解吧！

 */
import (
	"container/heap"
)

type Item struct {
	value    byte
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

//
//func (pq *PriorityQueue) update(item *Item, value byte, priority int) {
//	item.value = value
//	item.priority = priority
//	heap.Fix(pq, item.index)
//}

func leastInterval(tasks []byte, n int) int {

	dict := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		dict[tasks[i]]++
	}

	pq := make(PriorityQueue, len(dict))

	i := 0
	for value, priority := range dict {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	//
	ans := 0
	for pq.Len() > 0 {
		// need to pop n+1 items
		updates := make([]*Item, 0)
		for i = 0; i <= n; i++ {
			if pq.Len() == 0 {
				break
			}
			item := heap.Pop(&pq).(*Item)
			ans++
			if item.priority > 1 {
				item.priority -= 1
				updates = append(updates, item)
			}
		}

		if len(updates) > 0 {
			ans += n - i + 1 // 补全 n+1 个item, 如果 priority queue 里没有足够多的 top(n+1) item
			for j := 0; j < len(updates); j++ {
				heap.Push(&pq, updates[j])
			}
			updates = updates[:0]
		}
	}
	return ans
}
