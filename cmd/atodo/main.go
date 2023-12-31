package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	alphatodo "github.com/AlphaCoder147/ToDoAlpha"
)

const(
	todoFile = ".tododata.json"
)

func main(){
	add := flag.Bool("add", false, "Add anew Task")
	complete := flag.Int("complete", 0, "Mark a Task Complete")
	del := flag.Int("del", 0, "Delete a Task")
	list := flag.Bool("list", false, "List out tasks")

	flag.Parse()
	todos := &alphatodo.Todos{}
	if err := todos.Load(todoFile); err != nil{
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} 

	switch{
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.Add(task)
		err = todos.Store(todoFile)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.Print()

	default:
		fmt.Fprintln(os.Stdout, "Invalid command")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error){
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil{
		return " ", err
	}
	text := scanner.Text()
	if len(text) == 0{
		return "", errors.New("task cannot be empty")
	}
	return text, nil
}