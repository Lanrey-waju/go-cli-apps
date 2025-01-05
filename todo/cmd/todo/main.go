package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lanrey-waju/todo"
)

func main() {
	task := os.Args[1:]
	taskString := strings.Join(task, " ")
	il := todo.ItemList{}
	il.Add(taskString)

	fmt.Println(il[0].Task)
}
