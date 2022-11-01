package tasks

import (
	"testing"

	filet "github.com/Flaque/filet"
)

func TestRmFile(t *testing.T) {
	defer filet.CleanUp(t)

	tmp_file := filet.TmpFile(t, "", "").Name()
	args := RmFile{"path": tmp_file}

	args.Execute()
	if filet.Exists(t, tmp_file) {
		t.Error("Expect file  to be deleted")
	}
}

func TestRmFileCheckingForPathArgumentPresence(t *testing.T) {
	args := RmFile{"xx": "yy"}

	err := args.Execute()
	if err == nil {
		t.Error("Expect error to be returned if path argument is not specified")
	}
}
