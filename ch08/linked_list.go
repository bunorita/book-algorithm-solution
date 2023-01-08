package ch08

// Bidirectional Linked List (Doubly linked list)
type LinkedNode struct {
	key, value string
	next, prev *LinkedNode
}

func NewLinkedListNode(key, value string) *LinkedNode {
	return &LinkedNode{key: key, value: value}
}

type LinkedList struct {
	sentinel *LinkedNode
}

func NewLinkedList(values ...string) *LinkedList {
	dl := &LinkedList{
		sentinel: &LinkedNode{},
	}
	for i := len(values) - 1; i >= 0; i-- {
		dl.Prepend(&LinkedNode{value: values[i]})
	}
	return dl
}

// insert v after p
func (dl *LinkedList) Insert(v, p *LinkedNode) {
	if p == nil {
		p = dl.sentinel
	}
	// v - p.next
	v.next = p.next
	if p.next != nil {
		p.next.prev = v
	}

	// p - v
	p.next = v
	v.prev = p
}

func (dl *LinkedList) Prepend(v *LinkedNode) {
	dl.Insert(v, nil)
}

func (dl *LinkedList) Append(v *LinkedNode) {
	dl.Insert(v, dl.getLastNode())
}

func (dl *LinkedList) Erase(v *LinkedNode) {
	if v == nil {
		return
	}
	// O(1)
	if v.prev != nil {
		v.prev.next = v.next
	}
	if v.next != nil {
		v.next.prev = v.prev
	}
}

func (dl *LinkedList) GetNode(i int) *LinkedNode {
	if i < 0 {
		return nil
	}

	index := 0
	for cur := dl.sentinel.next; cur != nil; cur = cur.next {
		if index == i {
			return cur
		}
		index++
	}
	return nil
}

func (dl *LinkedList) GetNodeByValue(val string) *LinkedNode {
	for cur := dl.sentinel.next; cur != nil; cur = cur.next {
		if cur.value == val {
			return cur
		}
	}
	return nil
}

func (dl *LinkedList) getNodeByKey(key string) *LinkedNode {
	for cur := dl.sentinel.next; cur != nil; cur = cur.next {
		if cur.key == key {
			return cur
		}
	}
	return nil
}

func (dl *LinkedList) getLastNode() *LinkedNode {
	for cur := dl.sentinel.next; cur != nil; cur = cur.next {
		if cur.next == nil {
			return cur
		}
	}
	return nil
}

func (dl *LinkedList) String() string {
	var s string
	for cur := dl.sentinel.next; cur != nil; cur = cur.next {
		s += cur.value + " -> "
	}
	return s
}
