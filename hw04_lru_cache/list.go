package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len       int
	frontItem *ListItem
	backItem  *ListItem
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.frontItem
}

func (l list) Back() *ListItem {
	return l.backItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newFrontItem := ListItem{
		Value: v,
		Next:  l.frontItem,
		Prev:  nil,
	}
	if l.frontItem != nil {
		l.frontItem.Prev = &newFrontItem
	}
	l.frontItem = &newFrontItem
	if l.backItem == nil {
		l.backItem = &newFrontItem
	}
	l.len++

	return l.frontItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newBackItem := ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.backItem,
	}
	if l.backItem != nil {
		l.backItem.Next = &newBackItem
	}
	if l.frontItem == nil {
		l.frontItem = &newBackItem
	}
	l.backItem = &newBackItem
	l.len++

	return l.backItem
}

func (l *list) Remove(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if l.frontItem == i {
		l.frontItem = i.Next
	}
	if l.backItem == i {
		l.backItem = i.Prev
	}
	i = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i.Value)
	l.Remove(i)
}

func NewList() List {
	l := new(list)
	l.len = 0
	l.frontItem, l.backItem = nil, nil

	return l
}
