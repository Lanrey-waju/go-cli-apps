package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lanrey-waju/todo"
)

func main() {
	// parse command line flags
	task := flag.String("task", "", "add task to ToDo list")
	list := flag.Bool("list", false, "lists all tasks")
	complete := flag.Int("complete", 0, "item to be completed")

	flag.Parse()
	const todoFileName = "tasksFile.json"

	il := &todo.ItemList{}
	if err := il.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	case *list:
		// print list of tasks
		fmt.Print(il)
	case *task != "":
		il.Add(*task)
		if err := il.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete > 0:
		if err := il.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := il.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)
	}
}
