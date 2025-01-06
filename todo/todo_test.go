package todo_test

import (
	"os"
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

func TestSaveGet(t *testing.T) {
	il1 := todo.ItemList{}
	il2 := todo.ItemList{}

	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}

	for _, v := range tasks {
		il1.Add(v)
	}

	if il1[0].Task != tasks[0] {
		t.Errorf("expected %v, got %v", tasks[0], il1[0].Task)
	}

	file, err := os.CreateTemp("", "tempFile-")
	if err != nil {
		t.Fatalf("unable to create file: %v", err)
	}

	if err := il1.Save(file.Name()); err != nil {
		t.Fatalf("unable to save file into given file: %v", err)
	}

	if err := il2.Get(file.Name()); err != nil {
		t.Fatalf("unable to retrieve tasks from file: %v", err)
	}

	if len(il1) != len(il2) {
		t.Errorf("length of both item lists should be equal!")
	}

	if il2[0].Task != tasks[0] {
		t.Errorf("expected %v, got %v", tasks[0], il2[0].Task)
	}
}
