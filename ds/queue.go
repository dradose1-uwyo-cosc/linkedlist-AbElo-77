package ds

import (
	"errors"
)

type Queue struct {
    list LinkedList
}

func (q *Queue) Push(value string) {
	q.list.InsertAt(q.list.Size, value)
}

func (q *Queue) Pop() (string, error) {

	if q.list.Head == nil {
		return "", errors.New("Nothing to pop.")
	}

	val := q.list.Head.data
	q.list.RemoveAt(0)

	return val, nil
}