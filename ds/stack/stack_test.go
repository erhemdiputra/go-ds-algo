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

	for i := 1; i <= 4; i++ {
		stack.Push(i)
	}
	assert.Equal(t, 4, stack.Len())

	value = stack.Pop()
	assert.Equal(t, 4, value)

	value = stack.Peek()
	assert.Equal(t, 3, value)
}
