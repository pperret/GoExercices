package github

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

// getText gets the comment body using the prefered editor
func getText() (string, error) {
	// Gets the prefered editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		if runtime.GOOS == "windows" {
			editor = "notepad.exe"
		} else {
			editor = "vi"
		}
	}

	// Gets the full editor pathname
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		return "", err
	}

	// Creates a temporary file
	file, err := ioutil.TempFile("", "issue")
	if err != nil {
		return "", err
	}
	file.Close()
	filename := file.Name()

	// Runs the editor
	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, filename},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		os.Remove(filename)
		return "", err
	}

	// Gets the content of the temporary file
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Remove(filename)
		return "", err
	}
	body := string(bytes)

	// Removes the temporary file
	os.Remove(filename)
	return body, nil
}
