package tasks

import (
	"errors"
	"io/ioutil"
	"log"
	"testing"
)

type TestTask1 map[string]string
type TestTask2 map[string]string
type TestTask3 map[string]string
type TestTask4 map[string]string

var executionLog []string

func (_args TestTask1) Execute() error {
	executionLog = append(executionLog, "test_task1")
	return nil
}

func (_args TestTask2) Execute() error {
	executionLog = append(executionLog, "test_task2")
	return nil
}

func (_args TestTask3) Execute() error {
	executionLog = append(executionLog, "test_task3")
	return errors.New("you fail me")
}

func (_args TestTask4) Execute() error {
	executionLog = append(executionLog, "test_task4")
	return nil
}

func testTask1Executor(args map[string]string) task {
	return TestTask1(args)
}

func testTask2Executor(args map[string]string) task {
	return TestTask2(args)
}

func testTask3Executor(args map[string]string) task {
	return TestTask3(args)
}

func testTask4Executor(args map[string]string) task {
	return TestTask4(args)
}

func cleanUpExecutionLog() {
	executionLog = make([]string, 0)
}

func init() {
	log.SetOutput(ioutil.Discard)
	Register("test_task1", testTask1Executor)
	Register("test_task2", testTask2Executor)
	Register("test_task3", testTask3Executor)
	Register("test_task4", testTask4Executor)
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestExecutor(t *testing.T) {
	defer cleanUpExecutionLog()

	tasks := []Definition{
		Definition{Type: "test_task1"},
		Definition{Type: "test_task3"},
		Definition{Type: "test_task4"},
		Definition{Type: "test_task2"},
	}
	correct_order := []string{"test_task1", "test_task3", "test_task4", "test_task2"}

	Execute(tasks)

	if !Equal(executionLog, correct_order) {
		t.Errorf("Expect tasks to be executed in correct order, want = %v, got = %v", correct_order, executionLog)
	}
}

func TestExecutorWithSkip(t *testing.T) {
	defer cleanUpExecutionLog()

	tasks := []Definition{
		Definition{Type: "test_task1"},
		Definition{Type: "test_task3", Skip: true},
		Definition{Type: "test_task4"},
		Definition{Type: "test_task2"},
	}
	correct_order := []string{"test_task1", "test_task4", "test_task2"}

	Execute(tasks)

	if !Equal(executionLog, correct_order) {
		t.Errorf("Expect tasks to be executed in correct order, want = %v, got = %v", correct_order, executionLog)
	}
}

func TestExecutorWithAbortOnFail(t *testing.T) {
	defer cleanUpExecutionLog()

	tasks := []Definition{
		Definition{Type: "test_task1"},
		Definition{Type: "test_task3", AbortOnFail: true},
		Definition{Type: "test_task4"},
		Definition{Type: "test_task2"},
	}
	correct_order := []string{"test_task1", "test_task3"}

	Execute(tasks)

	if !Equal(executionLog, correct_order) {
		t.Errorf("Expect tasks to be executed in correct order, want = %v, got = %v", correct_order, executionLog)
	}
}
