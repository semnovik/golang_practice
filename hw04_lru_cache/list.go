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
	head   *ListItem
	tail   *ListItem
	length int
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case i.Prev == nil:
		break
	case i.Prev != nil && i.Next != nil:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		i.Next = l.Front()
		l.Front().Prev = i
		i.Prev = nil
		l.head = i
	case i.Prev != nil && i.Next == nil:
		i.Prev.Next = nil
		l.Front().Prev = i
		i.Prev = nil
		i.Next = l.Front()
		l.head = i
	}
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Prev == nil:
		l.head = i.Next
	case i.Next == nil:
		l.tail = i.Prev
	case i.Prev != nil:
		i.Prev.Next = i.Next
	case i.Next != nil:
		i.Next.Prev = i.Prev
	}

	i.Next = nil
	i.Prev = nil
	i.Value = nil

	l.length--
}

func (l *list) PushFront(v interface{}) *ListItem {
	node := &ListItem{Value: v}
	switch {
	case l.head == nil:
		l.head = node
	case l.head != nil && l.tail == nil:
		cache := l.head
		l.tail = cache
		l.tail.Prev = l.head

		l.head = node
		l.head.Next = l.tail
	case l.head != nil && l.tail != nil:
		l.head.Prev = node
		cacheHead := l.head
		l.head = node
		l.head.Next = cacheHead

	}
	l.length++
	return node
}

func (l *list) PushBack(v interface{}) *ListItem {
	node := &ListItem{Value: v}
	switch {
	case l.head == nil && l.tail == nil:
		l.head = node
	case l.tail == nil && l.head != nil:
		l.tail = node
		l.tail.Prev = l.head
		l.head.Next = node
	case l.head != nil && l.tail != nil:
		l.tail.Next = node
		cachedTail := l.tail
		l.tail = node
		l.tail.Prev = cachedTail
	}
	l.length++
	return node
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) Len() int {
	return l.length
}

func NewList() List {
	return new(list)
}
