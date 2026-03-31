package cli

import (
	"fmt"
	"os"
	"task-manager/internal/task"
	"flag"
)

func Run(store *task.Store) {
	command := os.Args[1]

	switch command {
	case "add":
		runAdd(store)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %q\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func runAdd(store *task.Store){
	fs:=flag.NewFlagSet("add",flag.ExitOnError)
	title:=fs.String("title","","title name(required)")
	fs.Parse(os.Args[2:])

	if *title==""{
		fmt.Fprintln(os.Stderr,"error:- title is required")
		fs.Usage()
		os.Exit(1)
	}

	t,err:= store.Add(*title)
	if err != nil{
		handleError(err)
		return
	}
	fmt.Printf("added: %s\n", t)
}



func printUsage() {
	fmt.Println(`task - a simple CLI task manager

usage:
	task add -title "buy groceries"
	task list
	`)
}
