package leetcode

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	h map[int]int
	// l []int  // 可以保存一个 list, 在 Remove 的时候，让 pos 和最后一个 element swap 然后更新  h 里面的 index
}

func Constructor() RandomizedSet {
	rand.Seed(int64(time.Now().UTC().Nanosecond()))
	return RandomizedSet{
		h: make(map[int]int),
	}

}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.h[val]; ok {
		// exist
		return false
	}
	this.h[val] = 1
	return true

}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.h[val]; !ok {
		return false
	}
	delete(this.h, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	n := len(this.h)
	pos := rand.Intn(n)
	i := 0

	for k := range this.h {
		if i == pos {
			return k
		}
		i++
	}
	return -1
}
