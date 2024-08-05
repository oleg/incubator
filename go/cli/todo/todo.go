package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string, now time.Time) {
	*l = append(*l, item{
		Task:        task,
		Done:        false,
		CreatedAt:   now,
		CompletedAt: time.Time{},
	})
}

func (l *List) Complete(i int, now time.Time) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = now
	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	*l = slices.Delete(ls, i-1, i)
	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}

func (l *List) ExtendedSting() string {
	formatted := ""
	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}
		created := t.CreatedAt.Format(time.RFC3339)
		completed := ""
		if !t.CompletedAt.IsZero() {
			completed = " - " + t.CompletedAt.Format(time.RFC3339)
		}
		formatted += fmt.Sprintf("%s%d: %s %s%s\n",
			prefix, k+1, t.Task, created, completed)
	}
	return formatted
}
