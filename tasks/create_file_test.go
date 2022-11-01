package tasks

import (
	"path/filepath"
	"testing"

	filet "github.com/Flaque/filet"
)

func TestCreateFile(t *testing.T) {
	defer filet.CleanUp(t)

	tmp_dir := filet.TmpDir(t, "")
	args := CreateFile{"path": filepath.Join(tmp_dir, "test_file")}

	args.Execute()
	if !filet.DirContains(t, tmp_dir, "test_file") {
		t.Error("Expect 'test_file' directory to be created")
	}
}

func TestCreateFileCheckingForPathArgumentPresence(t *testing.T) {
	args := CreateFile{"xx": "yy"}

	err := args.Execute()
	if err == nil {
		t.Error("Expect error to be returned if path argument is not specified")
	}
}
