package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	var value interface{}

	queue := New()
	assert.Equal(t, 0, queue.Len())
	assert.Nil(t, queue.Pop())
	assert.Nil(t, queue.Peek())

	for i := 1; i <= 4; i++ {
		queue.Push(i)
	}
	assert.Equal(t, 4, queue.Len())

	value = queue.Pop()
	assert.Equal(t, 1, value)

	value = queue.Peek()
	assert.Equal(t, 2, value)
}
