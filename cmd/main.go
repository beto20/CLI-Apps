package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/beto20/cli/todo"
)

const (
	todoFile = "data/todos.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	delete := flag.Int("delete", 0, "mark todo as deleted")
	list := flag.Bool("list", false, "list todos")
	addV2 := flag.Bool("task", false, "add title and description split with ','")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		addTask(todos)

	case *complete > 0:
		completeTask(todos, complete)

	case *delete > 0:
		deleteTask(todos, delete)

	case *list:
		todos.Print()

	case *addV2:
		addTaskV2(todos)

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}

func addTaskV2(todos *todo.Todos) {
	data, err := getInput(os.Stdin, flag.Args()...)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	input := strings.Split(data, ",")

	itemTemp := todo.ItemTemp{
		Title:       input[0],
		Description: input[1],
	}

	todos.Addv2(itemTemp)

	err = todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func addTask(todos *todo.Todos) {
	task, err := getInput(os.Stdin, flag.Args()...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	todos.Add(task)
	err = todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func completeTask(todos *todo.Todos, complete *int) {
	err := todos.Complete(*complete)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func deleteTask(todos *todo.Todos, delete *int) {
	err := todos.Delete(*delete)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil
}
