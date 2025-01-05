package todo_test

import (
	"testing"

	"github.com/lanrey-waju/todo"
)

func TestAdd(t *testing.T) {
	tl := todo.ItemList{}
	taskName := "code"
	tl.Add(taskName)

	if tl[0].Task != taskName {
		t.Errorf("expected %s, got %s", taskName, tl[0].Task)
	}
}

func TestComplete(t *testing.T) {
	tl := todo.ItemList{}
	taskName := "code"
	tl.Add(taskName)

	if tl[0].Task != taskName {
		t.Errorf("expected %s, got %s", taskName, tl[0].Task)
	}
	if tl[0].Done {
		t.Errorf("%s should not be done", tl[0].Task)
	}
	tl.Complete(1)

	if !tl[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	tl := todo.ItemList{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
		"New Task 4",
	}
	for _, task := range tasks {
		tl.Add(task)
	}

	if len(tl) != len(tasks) {
		t.Errorf("expected list length of %d, got %d", 4, len(tl))
	}

	tl.Delete(2)
	tl.Delete(3)

	if len(tl) == len(tasks) {
		t.Errorf("tasks should be %d, and not %d", len(tl), len(tasks))
	}
}
