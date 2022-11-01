package tasks

import (
	"errors"
	"os"
	"strconv"
)

type RmDir map[string]string

func init() {
	Register("rm_dir", rmDirExecutor)
}

func (args RmDir) validate() error {
	_, set := args["path"]
	if !set {
		return errors.New("path is not specified")
	}

	_, set = args["recursive"]
	if set {
		_, err := strconv.ParseBool(args["recursive"])

		return err
	}

	return nil
}

func (args RmDir) path() string {
	return args["path"]
}

func (args RmDir) recursive() bool {
	recursive, set := args["recursive"]
	if !set {
		return false
	}

	value, _ := strconv.ParseBool(recursive)

	return value
}

func (args RmDir) Execute() error {
	err := args.validate()
	if err != nil {
		return err
	}

	if args.recursive() {
		err = os.RemoveAll(args.path())
	} else {
		err = os.Remove(args.path())
	}
	return err
}

func rmDirExecutor(args map[string]string) task {
	return RmDir(args)
}
