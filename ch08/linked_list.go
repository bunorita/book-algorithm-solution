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
