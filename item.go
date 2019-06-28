package list

import (
	"errors"
	"fmt"
)
//Item структура "элемент списка"
type Item struct {
	value interface{}
	next  *Item
	prev  *Item
}

func newItem(value interface{}) *Item {
	return &Item{value: value}
}
// Value возвращаем значение или ошибку, если не проинициализировано еще
func (i *Item) Value() (interface{}, error) {
	if i != nil {
		return i.value, nil
	}
	return nil, errors.New("Value is nil")
}
// String реализуем интерфейс stringer, чтобы удобно было помотреть результаты
func (i *Item) String() string {
	if i != nil {
		return fmt.Sprintf("%v", i.value)
	}
	return ""
}
// Next возвращаем следующий элемент
func (i *Item) Next() *Item {
	return i.next
}
// Prev возвращаем предыдущий элемент
func (i *Item) Prev() *Item {
	return i.prev
}