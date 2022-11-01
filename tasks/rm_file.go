package tasks

import (
	"errors"
	"os"
)


func init() {
	Register("rm_file", rmFileExecutor)
}

type RmFile map[string]string

func (args RmFile) validate() error {
	_, set := args["path"]
	if !set {
		return errors.New("path is not specified")
	}

	return nil
}

func (args RmFile) path() string {
	return args["path"]
}

func (args RmFile) Execute() error {
	err := args.validate()
	if err != nil {
		return err
	}
	err = os.Remove(args.path())
	return nil
}

func rmFileExecutor(args map[string]string) task {
	return RmFile(args)
}
