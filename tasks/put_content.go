package tasks

import (
	"errors"
	"os"
	"strconv"
)

type PutContent map[string]string

func init() {
	Register("put_content", putContentExecutor)
}

func (args PutContent) validate() error {
	_, set := args["path"]
	if !set {
		return errors.New("path is not specified")
	}

	_, set = args["content"]
	if !set {
		return errors.New("content is not specified")
	}

	_, set = args["append"]
	if set {
		_, err := strconv.ParseBool(args["append"])

		if err != nil {
			return err
		}
	}

	return nil
}

func (args PutContent) path() string {
	return args["path"]
}

func (args PutContent) content() string {
	return args["content"]
}

func (args PutContent) append_content() bool {
	append_content, set := args["append"]
	if set {
		value, err := strconv.ParseBool(append_content)

		if err == nil {
			return value
		}
	}

	return false
}

func (args PutContent) Execute() error {
	err := args.validate()
	if err != nil {
		return err
	}

	flags := os.O_CREATE | os.O_WRONLY
	if args.append_content() {
		flags = flags | os.O_APPEND
	}

	file, err := os.OpenFile(args.path(), flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(args.content())

	return err
}

func putContentExecutor(args map[string]string) task {
	return PutContent(args)
}
