package ch08

func NewNode(name string) *Node {
	return &Node{name: name}
}

type Node struct {
	next *Node
	name string
}

type LinkedList struct {
	sentinel *Node
}

func NewList(names ...string) *LinkedList {
	l := &LinkedList{
		sentinel: NewNode(""), // sentinel node
	}
	for i := len(names) - 1; i >= 0; i-- {
		l.Prepend(NewNode(names[i]))
	}
	return l
}

// Insert v after p
func (l *LinkedList) Insert(v, p *Node) {
	if p == nil {
		p = l.sentinel
	}
	// p - v
	v.next = p.next
	p.next = v
}

func (l *LinkedList) Prepend(v *Node) {
	l.Insert(v, nil)
}

func (l *LinkedList) GetNodeByName(name string) *Node {
	// O(N)
	for cur := l.sentinel.next; cur != nil; cur = cur.next {
		if cur.name == name {
			return cur
		}
	}
	return nil
}

func (l *LinkedList) String() string {
	var s string
	for cur := l.sentinel.next; cur != nil; cur = cur.next {
		s += cur.name + " -> "
	}
	return s
}

// Bidirectional Linked List
type DoublyLinkedNode struct {
	name       string
	next, prev *DoublyLinkedNode
}

type DoublyLinkedList struct {
	sentinel *DoublyLinkedNode
}

func NewDoublyLinkedList(names ...string) *DoublyLinkedList {
	dl := &DoublyLinkedList{
		sentinel: &DoublyLinkedNode{},
	}
	for i := len(names) - 1; i >= 0; i-- {
		dl.Prepend(&DoublyLinkedNode{name: names[i]})
	}
	return dl
}

// insert v after p
func (dl *DoublyLinkedList) Insert(v, p *DoublyLinkedNode) {
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

func (dl *DoublyLinkedList) Prepend(v *DoublyLinkedNode) {
	dl.Insert(v, nil)
}

func (dl *DoublyLinkedList) Erase(v *DoublyLinkedNode) {
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

func (dl *DoublyLinkedList) GetNodeByName(name string) *DoublyLinkedNode {
	for cur := dl.sentinel.next; cur != nil; cur = cur.next {
		if cur.name == name {
			return cur
		}
	}
	return nil
}

func (dl *DoublyLinkedList) String() string {
	var s string
	for cur := dl.sentinel.next; cur != nil; cur = cur.next {
		s += cur.name + " -> "
	}
	return s
}
