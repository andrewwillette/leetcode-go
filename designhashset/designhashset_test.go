package designhashset

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/design-hashset/
type MyHashSet struct {
	set []uint32
}

func Constructor() MyHashSet {
	return MyHashSet{set: make([]uint32, 0)}
}

func (this *MyHashSet) Add(key int) {
	hash := hashInt(key)
	if len(this.set) == 0 {
		this.set = append(this.set, hash)
		return
	}
	for i, iter := range this.set {
		if iter == hash {
			return
		}
		if iter > hash {
			this.set = append(this.set, 0)
			copy(this.set[i+1:], this.set[i:])
			this.set[i] = hash
			return
		}
	}
	this.set = append(this.set, hash)
}

func (this *MyHashSet) Remove(key int) {
	hash := hashInt(key)
	for i, iter := range this.set {
		if iter > hash {
			return
		}
		if iter == hash {
			this.set = append(this.set[:i], this.set[i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	hash := hashInt(key)
	for _, iter := range this.set {
		if iter == hash {
			return true
		}
		if iter > hash {
			return false
		}
	}
	return false
}

func hashInt(toHash int) uint32 {
	h := fnv.New32a()
	h.Write([]byte(strconv.Itoa(toHash)))
	return h.Sum32()
}

func TestHashInt(t *testing.T) {
	h1 := hashInt(8)
	h2 := hashInt(9)
	h3 := hashInt(8)
	fmt.Printf("%v\n", h1 < h2)
	fmt.Printf("%v\n", h2 < h1)
	fmt.Printf("%v\n", h3 == h1)
}

func TestMyHashSet(t *testing.T) {
	hashset := Constructor()
	hashset.Add(1)
	hashset.Add(2)
	fmt.Printf("%v\n", len(hashset.set))
	require.True(t, hashset.Contains(1))
}
