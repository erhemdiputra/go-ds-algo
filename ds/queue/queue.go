package queue

import (
	"container/list"
)

type Queue struct {
	list *list.List
}

func New() *Queue {
	q := Queue{
		list: list.New(),
	}
	return &q
}

func (q *Queue) Push(item interface{}) {
	q.list.PushBack(item)
}

func (q *Queue) Pop() interface{} {
	if q.Len() > 0 {
		e := q.list.Front()
		q.list.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) Peek() interface{} {
	if q.Len() > 0 {
		return q.list.Front().Value
	}
	return nil
}

func (q *Queue) Len() int {
	return q.list.Len()
}
