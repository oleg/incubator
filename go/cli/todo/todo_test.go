package todo_test

import (
	"github.com/oleg/incubator/go/cli/todo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Add(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	assert.Equal(t, taskName, l[0].Task)
}
