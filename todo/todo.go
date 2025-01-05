package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string    `json:"task"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

type ItemList []item

// Add adds a new task to the list
func (il *ItemList) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*il = append(*il, t)
}

// Complete marks the item of the given number ad done or completed
func (il *ItemList) Complete(i int) error {
	ls := *il
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete removes the item from the task list and rearranges the list
func (il *ItemList) Delete(i int) error {
	ls := *il
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*il = append(ls[:i-1], ls[i:]...)

	return nil
}

func (il *ItemList) Save(filename string) error {
	dat, err := json.Marshal(il)
	if err != nil {
		return fmt.Errorf("unable to marshal into json: %v", err)
	}

	return os.WriteFile(filename, dat, 0o644)
}

func (il *ItemList) Get(filename string) error {
	dat, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("unable to read file: %v", err)
	}
	if len(dat) == 0 {
		return nil
	}

	return json.Unmarshal(dat, il)
}