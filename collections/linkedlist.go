package collections

import (
	"reflect"
)

type llitem struct {
	item interface{}
	next *llitem
}

type LinkedList struct {
	first *llitem
	iter  *llitem
}

// Get gets the item at the specified index.
func (l *LinkedList) Get(index int) (reflect.Type, interface{}) {
	var it = l.first
	if it != nil {
		for i := 1; i < index+1; i++ {
			it = it.next
			if it == nil {
				return nil, nil
			}
		}
		return reflect.TypeOf(it.item), it.item
	}
	return nil, nil
}

// Insert inserts an item at the specified index.
func (l *LinkedList) Insert(index int, item interface{}) {
	var it = l.first
	for i := 1; i != index-1; i++ {
		it = it.next
	}
	it.next = &llitem{
		item: item,
		next: it.next.next,
	}
}

// PushBack pushes an item to the end of the list.
func (l *LinkedList) PushBack(item interface{}) {
	if l.first != nil {
		var it = l.first
		for i := 1; it.next != nil; i++ {
			it = it.next
		}
		it.next = &llitem{
			item: item,
			next: nil,
		}
		return;
	}
	l.first = &llitem{
		item: item,
		next: nil,
	}
}

// PushFront pushes an item to the front of the list.
func (l *LinkedList) PushFront(item interface{}) {
	l.first = &llitem{
		item: item,
		next: l.first,
	}
}

// Length gets the length of the list.
func (l *LinkedList) Length() int {
	var it = l.first
	if it != nil {
		i := 1
		for it.next != nil {
			it = it.next
			i++
		}
		return i
	}
	return 0
}

// Remove removes an item at the specified index.
func (l *LinkedList) Remove(index int) {
	if l.first != nil {
		if index > 0 {
			var it = l.first
			for i := 1; it.next != nil && i != index-1; i++ {
				it = it.next
			}
			if it.next != nil {
				it.next = it.next.next
			}
		}
		nxt := l.first.next
		l.first = nil
		l.first = nxt
	}
}

// Next advances the iterator to the next item, returning the item.
// Next will revert back to the front item if the next item is null.
func (l *LinkedList) Next() interface{} {
	if l.iter == nil {
		l.Reset()
		return l.iter.item
	}

	l.iter = l.iter.next
	if (l.iter != nil) {
		return l.iter.item
	}
	return nil
}

// Current gets the current item in the iterator.
// Current also returns the type of the item.
func (l *LinkedList) Current() (vtype reflect.Type, value interface{}) {
	if l.iter != nil {
		return reflect.TypeOf(l.iter.item), l.iter.item
	}
	return nil, nil
}

// Reset resets the iterator.
func (l *LinkedList) Reset() {
	l.iter = l.first
}

// NewLinkedList constructs a new generic linked list.
func NewLinkedList(item interface{}) *LinkedList {
	first := &llitem{
		item: item,
		next: nil,
	}
	return &LinkedList{
		first: first,
	}
}
