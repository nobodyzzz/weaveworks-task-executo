package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"weaveworks-task-executor/tasks"
)

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}

func setupUsage(f *flag.FlagSet) {
	f.Usage = func() {
		w := f.Output()

		fmt.Fprintf(w, "Usage of %s [options] <file1 file2 ... fileN>:\n", os.Args[0])

		flag.PrintDefaults()
	}
}

func main() {
	setupUsage(flag.CommandLine)

	silent := flag.Bool("silent", false, "Silence log output")

	flag.Parse()

	if *silent {
		log.SetOutput(ioutil.Discard)
	}

	for _, tasksDefinitionsFile := range flag.Args() {
		var tasksDefinitions []tasks.Definition

		log.SetPrefix(fmt.Sprintf("[%s] ", tasksDefinitionsFile))
		err := tasks.LoadFromFile(tasksDefinitionsFile, &tasksDefinitions)
		check(err)

		err = tasks.Execute(tasksDefinitions)
		check(err)
	}
}
