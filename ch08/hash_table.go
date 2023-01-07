package ch08

import (
	"fmt"
	"strings"
)

type HashTable struct {
	t []*LinkedList // values
	m int           // size of t
}

func NewHashTable(m int) *HashTable {
	t := make([]*LinkedList, m)
	for i := range t {
		t[i] = NewLinkedList()
	}
	return &HashTable{
		m: m,
		t: t,
	}
}

func (h *HashTable) PrintArray() {
	strs := make([]string, h.m)
	for i := range h.t {
		strs[i] = fmt.Sprintf("%d:%q", i, h.t[i])
	}
	fmt.Println("[" + strings.Join(strs, ", ") + "]")
}

func (h *HashTable) hash(x string) int {
	var sum int
	for _, char := range x {
		sum += int(char)
	}
	return sum % h.m // rolling hash
}

func (h *HashTable) Set(key, val string) {
	h.t[h.hash(key)].Append(NewLinkedListNode(key, val))
}

func (h *HashTable) Get(key string) string {
	node := h.t[h.hash(key)].getNodeByKey(key)
	if node == nil {
		return ""
	}
	return node.value
}
