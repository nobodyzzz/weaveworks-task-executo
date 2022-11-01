package tasks

import (
	"errors"
	"os"
	"strconv"
)

type CreateDir map[string]string

func init() {
	Register("create_dir", createDirExecutor)
}

func (args CreateDir) validate() error {
	_, set := args["path"]
	if !set {
		return errors.New("path is not specified")
	}

	_, set = args["create_intermediate"]
	if set {
		_, err := strconv.ParseBool(args["create_intermediate"])

		if err != nil {
			return err
		}
	}

	return nil
}

func (args CreateDir) path() string {
	return args["path"]
}

func (args CreateDir) create_intermediate() bool {
	create_intermediate, set := args["create_intermediate"]
	if set {
		value, err := strconv.ParseBool(create_intermediate)
		if err == nil {
			return value
		}
	}

	return false
}

func (args CreateDir) Execute() error {
	err := args.validate()
	if err != nil {
		return err
	}

	if args.create_intermediate() {
		err = os.MkdirAll(args.path(), 0750)
	} else {
		err = os.Mkdir(args.path(), 0750)
	}

	if err != nil && errors.Is(err, os.ErrExist) {
		return nil
	}

	return err
}

func createDirExecutor(args map[string]string) task {
	return CreateDir(args)
}
