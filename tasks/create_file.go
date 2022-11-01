package tasks

import (
	"errors"
	"os"
)

type CreateFile map[string]string

func init() {
	Register("create_file", createFileExecutor)
}

func (args CreateFile) validate() error {
	_, set := args["path"]
	if !set {
		return errors.New("path is not specified")
	}

	return nil
}

func (args CreateFile) path() string {
	return args["path"]
}

func (args CreateFile) Execute() error {
	err := args.validate()
	if err != nil {
		return err
	}

	return os.WriteFile(args.path(), []byte(""), 0644)
}

func createFileExecutor(args map[string]string) task {
	return CreateFile(args)
}
