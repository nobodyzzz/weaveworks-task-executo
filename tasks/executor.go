package tasks

import (
	"errors"
	"fmt"
	"log"
)

type task interface {
	Execute() error
}

var tasksExecutors map[string]func(map[string]string) task
var supportedTasks []string

func Execute(tasksDefinitions []Definition) error {
	for _, task := range tasksDefinitions {
		executor, found := tasksExecutors[task.Type]

		if task.Skip {
			log.Printf("SKIP task '%s'\n", task.Description())
			continue
		}

		if !found {
			log.Printf("Unknown task type '%s'. Supported tasks: %v", task.Type, supportedTasks)
			continue
		}

		log.Printf("START task '%s'\n", task.Description())
		err := executor(task.Args).Execute()
		if err != nil {
			msg := fmt.Sprintf("Task '%s' failed (%v)", task.Description(), err)
			if task.AbortOnFail {
				return errors.New(msg)
			}
			log.Println(msg)
		} else {
			log.Printf("END task '%s'\n", task.Description())
		}
	}
	return nil
}

func Register(taskName string, executorFunc func(map[string]string) task) {
	if tasksExecutors == nil {
		tasksExecutors = make(map[string]func(map[string]string) task)
	}

	_, already_registered := tasksExecutors[taskName]
	if already_registered {
		log.Printf("%s was registed more than once\n", taskName)
	}

	tasksExecutors[taskName] = executorFunc

	supportedTasks = append(supportedTasks, taskName)
}
