package todo_test

import (
	"github.com/oleg/incubator/go/cli/todo"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestList_Add(t *testing.T) {
	//given
	l := todo.List{}
	taskName := "New Task"

	//when
	l.Add(taskName, time.Now())

	//then
	assert.Equal(t, taskName, l[0].Task)
}

func TestList_Complete(t *testing.T) {
	//given
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName, time.Now())
	assert.False(t, l[0].Done)

	//when
	err := l.Complete(1, time.Now())

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
		l.Add(task, time.Now())
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
	l1.Add(taskName, time.Now())

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

func TestList_ExtendedSting(t *testing.T) {
	// given
	now := mustParse(t, "2021-07-01T00:00:00Z")
	l := todo.List{}
	l.Add("the first task", now.Add(time.Hour))
	l.Add("the second task", now.Add(2*time.Hour))
	l.Add("the third task", now.Add(3*time.Hour))
	err := l.Complete(2, now.Add(4*time.Hour))
	assert.NoError(t, err)

	// when
	str := l.ExtendedSting()

	// then
	assert.Equal(t, `  1: the first task 2021-07-01T01:00:00Z
X 2: the second task 2021-07-01T02:00:00Z - 2021-07-01T04:00:00Z
  3: the third task 2021-07-01T03:00:00Z
`, str)
}

func mustParse(t *testing.T, val string) time.Time {
	parse, err := time.Parse(time.RFC3339, val)
	assert.NoError(t, err)
	return parse
}
