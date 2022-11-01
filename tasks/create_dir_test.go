package tasks

import (
	"path/filepath"
	"testing"

	filet "github.com/Flaque/filet"
)

func TestCreateDir(t *testing.T) {
	defer filet.CleanUp(t)

	tmp_dir := filet.TmpDir(t, "")
	args := CreateDir{"path": filepath.Join(tmp_dir, "test_dir")}

	args.Execute()
	if !filet.DirContains(t, tmp_dir, "test_dir") {
		t.Error("Expect 'test_dir' directory to be created")
	}
}

func TestCreateDirWithIntermediate(t *testing.T) {
	defer filet.CleanUp(t)

	tmp_dir := filet.TmpDir(t, "")
	args := CreateDir{"path": filepath.Join(tmp_dir, "test_dir/subdir"), "create_intermediate": "true"}

	args.Execute()
	if !filet.DirContains(t, filepath.Join(tmp_dir, "test_dir"), "subdir") {
		t.Error("Expect 'test_dir' directory to be created with 'subdir' directory in it")
	}
}

func TestCreateDirCheckingForPathArgumentPresence(t *testing.T) {
	args := CreateDir{"create_intermediate": "true"}

	err := args.Execute()
	if err == nil {
		t.Error("Expect error to be returned if path argument is not specified")
	}
}
