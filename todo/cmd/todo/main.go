package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lanrey-waju/todo"
)

func main() {
	const fileName = "tasksFile.json"
	il := &todo.ItemList{}
	if err := il.Get(fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	case len(os.Args) == 1:
		for _, item := range *il {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		il.Add(item)
		il.Save(fileName)
	}
}
