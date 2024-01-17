package insertdeletegetrandom

import "math/rand"

// https://leetcode.com/problems/insert-delete-getrandom-o1/
type RandomizedSet struct {
	s map[int]bool
}

func Constructor() RandomizedSet {
	set := make(map[int]bool)
	return RandomizedSet{s: set}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.s[val]
	this.s[val] = true
	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.s[val]
	delete(this.s, val)
	return ok
}

func (this *RandomizedSet) GetRandom() int {
	keys := make([]int, len(this.s))
	i := 0
	for k := range this.s {
		keys[i] = k
		i++
	}
	return keys[rand.Intn(len(keys))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
