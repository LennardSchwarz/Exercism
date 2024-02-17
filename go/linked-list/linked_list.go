package linkedlist

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type List struct {
	head *Node
	tail *Node
}

type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

func NewList(elements ...interface{}) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	// check if the list is empty
	if l.head == nil {
		// create a new node
		newNode := &Node{Value: v, next: nil, prev: nil}
		// set the head and tail to the new node
		l.head = newNode
		l.tail = newNode
		return
	}
	// create a new node
	newNode := &Node{Value: v, next: l.head, prev: nil}
	l.head = newNode
	l.head.next.prev = newNode
}

func (l *List) Push(v interface{}) {
	// check if the list is empty
	if l.head == nil {
		// create a new node
		newNode := &Node{Value: v, next: nil, prev: nil}
		// set the head and tail to the new node
		l.head = newNode
		l.tail = newNode
		return
	}
	// range over the list to find the last node
	for n := l.head; n != nil; n = n.next {
		// if the next of the current node is nil, then the current node is the last node
		if n.next == nil {
			// create a new node
			newNode := &Node{Value: v, next: nil, prev: n}
			// set the next of the current node to the new node
			n.next = newNode
			// set the tail to the new node
			l.tail = newNode
			return
		}
	}
}

func (l *List) Shift() (interface{}, error) {
	// check if the list is empty
	if l.head == nil {
		return nil, nil
	}
	// check if the list has only one node
	if l.head.next == nil {
		toShift := l.head
		// empty the list
		l.head = nil
		l.tail = nil
		return toShift.Value, nil
	}
	toShift := l.head
	l.head = l.head.next
	l.head.prev = nil
	return toShift.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	// check if the list is empty
	if l.tail == nil {
		return nil, nil
	}
	// get prev of tail
	prev := l.tail.prev

	// remove the next of prev
	prev.next = nil

	return l.tail.Value, nil
}

func (l *List) Reverse() {
	current := l.head
	var temp *Node

	// Swap next and prev for all nodes of the list
	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev // move to next node which is previous due to swap
	}

	// Swap the head and tail pointers.
	if temp != nil {
		l.head, l.tail = l.tail, l.head
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
