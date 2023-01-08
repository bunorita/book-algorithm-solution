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
	size     int
}

func NewLinkedList(values ...string) *LinkedList {
	l := &LinkedList{
		sentinel: &LinkedNode{},
	}
	for i := len(values) - 1; i >= 0; i-- {
		l.Prepend(&LinkedNode{value: values[i]})
	}
	return l
}

func (l *LinkedList) Size() int {
	return l.size
}

// insert v after p
func (l *LinkedList) Insert(v, p *LinkedNode) {
	if p == nil {
		p = l.sentinel
	}
	// v - p.next
	v.next = p.next
	if p.next != nil {
		p.next.prev = v
	}

	// p - v
	p.next = v
	v.prev = p

	l.size++
}

func (l *LinkedList) Prepend(v *LinkedNode) {
	l.Insert(v, nil)
}

func (l *LinkedList) Append(v *LinkedNode) {
	l.Insert(v, l.getLastNode())
}

func (l *LinkedList) Erase(v *LinkedNode) {
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

	l.size--
}

func (l *LinkedList) GetNode(i int) *LinkedNode {
	if i < 0 {
		return nil
	}

	index := 0
	for cur := l.sentinel.next; cur != nil; cur = cur.next {
		if index == i {
			return cur
		}
		index++
	}
	return nil
}

func (l *LinkedList) GetNodeByValue(val string) *LinkedNode {
	for cur := l.sentinel.next; cur != nil; cur = cur.next {
		if cur.value == val {
			return cur
		}
	}
	return nil
}

func (l *LinkedList) getNodeByKey(key string) *LinkedNode {
	for cur := l.sentinel.next; cur != nil; cur = cur.next {
		if cur.key == key {
			return cur
		}
	}
	return nil
}

func (l *LinkedList) getLastNode() *LinkedNode {
	for cur := l.sentinel.next; cur != nil; cur = cur.next {
		if cur.next == nil {
			return cur
		}
	}
	return nil
}

func (l *LinkedList) String() string {
	var s string
	for cur := l.sentinel.next; cur != nil; cur = cur.next {
		s += cur.value + " -> "
	}
	return s
}
