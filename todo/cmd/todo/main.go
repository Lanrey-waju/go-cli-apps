package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lanrey-waju/todo"
)

// getTask takes an io.Reader and returns a list of
// task strings with a potential error
func getTask(r io.Reader) ([]string, error) {
	tasks := []string{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			return tasks, nil
		} else {
			tasks = append(tasks, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return tasks, nil
}

func main() {
	// parse command line flags
	add := flag.Bool(
		"add",
		false,
		"Add tasks to todo list. Run without arguments to enable multiline input from STDIN. Each line is a new task in the list. A blank line saves the list.",
	)
	list := flag.Bool("list", false, "lists all tasks")
	complete := flag.Int("complete", 0, "item to be completed")
	del := flag.Int("del", 0, "delete item from todo list")
	verbose := flag.Bool("verbose", false, "enable verbose output")

	flag.Parse()
	todoFileName := "tasksFile.json"

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	il := &todo.ItemList{}
	if err := il.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	case *list:
		// print list of tasks
		fmt.Print(il)
	case *add:
		if len(flag.Args()) > 0 {
			il.Add(strings.Join(flag.Args(), " "))

			if err := il.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			os.Exit(0)
		}
		tasks, err := getTask(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, task := range tasks {
			il.Add(task)
		}
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
	case *del > 0:
		if err := il.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := il.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *verbose:
		fmt.Print(il.PrintVerbose())
	default:
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)
	}
}
