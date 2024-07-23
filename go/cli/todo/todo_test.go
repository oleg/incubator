package todo_test

import (
	"github.com/oleg/incubator/go/cli/todo"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestList_Add(t *testing.T) {
	//given
	l := todo.List{}
	taskName := "New Task"

	//when
	l.Add(taskName)

	//then
	assert.Equal(t, taskName, l[0].Task)
}

func TestList_Complete(t *testing.T) {
	//given
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	assert.False(t, l[0].Done)

	//when
	err := l.Complete(1)

	//then
	assert.NoError(t, err)
	assert.True(t, l[0].Done)
}

func TestList_Delete(t *testing.T) {
	//given
	l := todo.List{}
	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}
	for _, task := range tasks {
		l.Add(task)
	}
	assert.Len(t, l, 3)

	//when
	err := l.Delete(2)

	//then
	assert.NoError(t, err)
	assert.Len(t, l, 2)
}

func TestList_Get(t *testing.T) {
	//given
	l1 := todo.List{}
	l2 := todo.List{}
	taskName := "New Task"
	l1.Add(taskName)

	temp, err := os.CreateTemp("", "")
	assert.NoError(t, err)
	defer os.Remove(temp.Name())

	//when
	err = l1.Save(temp.Name())
	assert.NoError(t, err)

	err = l2.Get(temp.Name())
	assert.NoError(t, err)

	//then
	assert.Equal(t, l1[0].Task, l2[0].Task)
}
