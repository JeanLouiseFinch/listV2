package list

import (
	"errors"
	"fmt"
	"sync"
)

// List структура "список", все поля неэкспортируемые
type List struct {
	first *Item
	last  *Item
	len   int
	sync.Mutex
}

//NewList создает новый список и возвращает указатель на него
func NewList() *List {
	return &List{}
}

//Len возвращает длину списка
func (l *List) Len() int {
	return l.len
}
// String реализуем интерфейс stringer, чтобы удобно было помотреть результаты
func (l *List) String() string {
	var result string
	result += fmt.Sprintf("Len: %d. ", l.len)
	for i := l.first; i != nil; i = i.next {
		if i == l.first {
			result += fmt.Sprintf("(%v)", i.value)
		} else {
			result += fmt.Sprintf("<--->(%v)", i.value)
		}
	}
	return result
}
// PushFront вставка нового элемента в начало
func (l *List) PushFront(value interface{}) {
	l.Lock()
	defer func() {
		l.len++
		l.Unlock()
	}()
	item := newItem(value)
	if l.first != nil {
		l.first.prev = item
		item.next = l.first
		l.first = item
	} else {
		l.first, l.last = item, item
	}
}
//PushBack вставка нового элемента в конец
func (l *List) PushBack(value interface{}) {
	l.Lock()
	defer func() {
		l.len++
		l.Unlock()
	}()
	item := newItem(value)
	if l.last != nil {
		l.last.next = item
		item.prev = l.last
		l.last = item
	} else {
		l.first, l.last = item, item
	}
}
// Remove удаляет элемент по ссылке из списка
func (l *List) Remove(item *Item) {
	// чтобы не было паники, если мы еще раз захотим удалить уже удаленный элемент
	if item.next == nil && item.prev == nil && item != l.first {
		return
	}
	l.Lock()
	defer func() {
		item.next, item.prev = nil, nil
		l.len--
		l.Unlock()
	}()
	switch {
	case item == l.first && l.first.next != nil: // случай, что элемент первый и не единственный
		l.first.next.prev = nil
		l.first = l.first.next
	case item == l.last && l.last.prev != nil: // случай, что элемент последний и не единственный
		l.last.prev.next = nil
		l.last = l.last.prev
	case item == l.first && item == l.last: // элемент единственный
		l.first, l.last = nil, nil
	default: // элемент внутри списка
		item.prev.next = item.next
		item.next.prev = item.prev
	}
}
//First возвращает ссылку на первый элемент или ошибку
func (l *List) First() (*Item, error) {
	if l.first != nil {
		return l.first, nil
	}
	return nil, errors.New("List empty")
}
// Last возвращает ссылку на последний элемент или ошибку
func (l *List) Last() (*Item, error) {
	if l.last != nil {
		return l.last, nil
	}
	return nil, errors.New("List empty")
}