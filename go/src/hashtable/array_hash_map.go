package hashtable

import "fmt"

type entry struct {
	key int
	val string
}

type arrayHashMap struct {
	buckets []*entry
}

func newArrayHashMap() *arrayHashMap {
	buckets := make([]*entry, 100)
	return &arrayHashMap{buckets: buckets}
}

func (a *arrayHashMap) hashFunc(key int) int {
	index := key % 100
	return index
}

func (a *arrayHashMap) get(key int) string {
	index := a.hashFunc(key)
	pair := a.buckets[index]
	if pair == nil {
		return "Not Found"
	}
	return pair.val
}

func (a *arrayHashMap) put(key int, val string) {
	pair := &entry{key: key, val: val}
	index := a.hashFunc(key)
	a.buckets[index] = pair
}

func (a *arrayHashMap) remove(key int) {
	index := a.hashFunc(key)
	a.buckets[index] = nil
}

func (a *arrayHashMap) entrySet() []*entry {
	var pairs []*entry
	for _, pair := range a.buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func (a *arrayHashMap) keySet() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

func (a *arrayHashMap) valueSet() []string {
	var values []string
	for _, pair := range a.buckets {
		if pair != nil {
			values = append(values, pair.val)
		}
	}
	return values
}

func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.val)
		}
	}
}
