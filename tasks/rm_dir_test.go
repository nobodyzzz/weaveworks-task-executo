package tasks

import (
	"path/filepath"
	"testing"

	filet "github.com/Flaque/filet"
)

func TestRmDir(t *testing.T) {
	defer filet.CleanUp(t)

	tmp_dir := filet.TmpDir(t, "")
	tmp_subdir := filepath.Join(tmp_dir, "test_subdir")
	CreateDir{"path": tmp_subdir}.Execute()

	args := RmDir{"path": tmp_subdir}

	args.Execute()
	if filet.DirContains(t, tmp_dir, "test_subdir") {
		t.Error("Expect 'test_subdir' directory to be deleted")
	}
}

func TestRmDirFailsToRemoveNonEmptyDir(t *testing.T) {
	defer filet.CleanUp(t)

	tmp_dir := filet.TmpDir(t, "")
	tmp_subdir := filepath.Join(tmp_dir, "test_subdir")
	CreateDir{"path": tmp_subdir}.Execute()
	CreateDir{"path": filepath.Join(tmp_subdir, "test_subdir1")}.Execute()
	CreateDir{"path": filepath.Join(tmp_subdir, "test_subdir2")}.Execute()
	CreateFile{"path": filepath.Join(tmp_subdir, "test_file")}.Execute()

	args := RmDir{"path": tmp_subdir}

	err := args.Execute()
	if err == nil {
		t.Error("Expect error to be returned when trying to delete non-empty directory without 'recursive' argument")
	}
}

func TestRmRemoveNonEmptyDirWithRecursive(t *testing.T) {
	defer filet.CleanUp(t)

	tmp_dir := filet.TmpDir(t, "")
	tmp_subdir := filepath.Join(tmp_dir, "test_subdir")
	CreateDir{"path": tmp_subdir}.Execute()
	CreateDir{"path": filepath.Join(tmp_subdir, "test_subdir1")}.Execute()
	CreateDir{"path": filepath.Join(tmp_subdir, "test_subdir2")}.Execute()
	CreateFile{"path": filepath.Join(tmp_subdir, "test_file")}.Execute()

	args := RmDir{"path": tmp_subdir, "recursive": "true"}

	args.Execute()
	if filet.DirContains(t, tmp_dir, "test_subdir") {
		t.Error("Expect 'test_subdir' directory to be deleted")
	}
}

func TestRmDirCheckingForPathArgumentPresence(t *testing.T) {
	args := RmDir{"recursive": "true"}

	err := args.Execute()
	if err == nil {
		t.Error("Expect error to be returned if path argument is not specified")
	}
}
