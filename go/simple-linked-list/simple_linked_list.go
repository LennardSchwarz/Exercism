package linkedlist

import "errors"

// Define the List and Element types here.
type List struct {
	head *Element
}

type Element struct {
	data int
	next *Element
}

func New(elements []int) *List {
	list := &List{}
	for _, element := range elements {
		list.Push(element)
	}
	return list
}

func (l *List) Size() int {
	size := 0
	for e := l.head; e != nil; e = e.next {
		size++
	}
	return size
}

func (l *List) Push(element int) {
	// add element at the end of the list
	newElement := &Element{data: element, next: nil}
	if l.head == nil {
		l.head = newElement
		return
	}
	for e := l.head; e != nil; e = e.next {
		if e.next == nil {
			e.next = newElement
			break
		}
	}
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("cannot pop empty list")
	}
	var prev = l.head
	for e := l.head; e != nil; e = e.next {
		if e.next == nil {
			// test if the list has only one element
			if prev == e {
				l.head = nil
			}
			prev.next = nil
			break
		}
		prev = e
	}
	return prev.data + 1, nil
}

func (l *List) Array() []int {
	var elements []int
	for e := l.head; e != nil; e = e.next {
		elements = append(elements, e.data)
	}
	return elements
}

func (l *List) Reverse() *List {
	// reverse the list
	var prev *Element
	current := l.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
	return l
}
