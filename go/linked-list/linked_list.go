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

// Insert value at the front of the list.
func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v, next: l.head, prev: nil}
	if l.head != nil {
		l.head.prev = newNode
	} else {
		l.tail = newNode // Set tail if list was empty
	}
	l.head = newNode
}

// Insert value at the back of the list.
func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v, next: nil, prev: l.tail}
	if l.tail != nil {
		l.tail.next = newNode
	} else {
		l.head = newNode // Set head if list was empty
	}
	l.tail = newNode
}

// Remove value from the front of the list.
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

// Remove value from the back of the list.
func (l *List) Pop() (interface{}, error) {
	// check if the list is empty
	if l.tail == nil {
		return nil, nil
	}
	toPop := l.tail
	if l.head == l.tail {
		// List has only one node
		l.head = nil
		l.tail = nil
	} else {
		// List has more than one node
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	return toPop.Value, nil

}

// Reverse the list.
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
	if l.head != nil {
		l.head, l.tail = l.tail, l.head
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
