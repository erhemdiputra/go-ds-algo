package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	var value interface{}

	stack := New()
	assert.Equal(t, 0, stack.Len())
	assert.Nil(t, stack.Pop())
	assert.Nil(t, stack.Peek())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	assert.Equal(t, 3, stack.Len())

	value = stack.Pop()
	assert.Equal(t, 3, value)

	value = stack.Peek()
	assert.Equal(t, 2, value)
}
