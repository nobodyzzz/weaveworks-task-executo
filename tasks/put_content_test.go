package tasks

import (
	"testing"

	filet "github.com/Flaque/filet"
)

func TestPutContent(t *testing.T) {
	defer filet.CleanUp(t)

	tmp_file := filet.TmpFile(t, "", "").Name()
	args := PutContent{"path": tmp_file, "content": "test content"}

	args.Execute()
	if !filet.FileSays(t, tmp_file, []byte("test content")) {
		t.Error("Expect file to have 'test content' in it")
	}
}

func TestPutContentWithAppend(t *testing.T) {
	defer filet.CleanUp(t)

	tmp_file := filet.TmpFile(t, "", "old content ").Name()
	args := PutContent{"path": tmp_file, "content": "appended content", "append": "true"}

	args.Execute()
	if !filet.FileSays(t, tmp_file, []byte("old content appended content")) {
		t.Error("Expect file to have old content with appended content in it")
	}
}

func TestPutContentCheckingForPathArgumentPresence(t *testing.T) {
	args := PutContent{"content": "appended content", "append": "true"}

	err := args.Execute()
	if err == nil {
		t.Error("Expect error to be returned if path argument is not specified")
	}
}

func TestPutContentCheckingForContentArgumentPresence(t *testing.T) {
	args := PutContent{"path": "path", "append": "true"}

	err := args.Execute()
	if err == nil {
		t.Error("Expect error to be returned if content argument is not specified")
	}
}
