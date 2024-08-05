package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tools %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCli(t *testing.T) {

	dir, err := os.Getwd()
	assert.NoError(t, err)

	cmdPath := filepath.Join(dir, binName)

	task1 := "test task number 1"
	t.Run("AddNewTaskFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task1)
		err := cmd.Run()
		assert.NoError(t, err)
	})
	task2 := "test task number 2"
	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdIn, err := cmd.StdinPipe()
		assert.NoError(t, err)
		_, err = io.WriteString(cmdStdIn, task2)
		assert.NoError(t, err)
		err = cmdStdIn.Close()
		assert.NoError(t, err)
		err = cmd.Run()
		assert.NoError(t, err)
	})
	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("  1: %s\n  2: %s\n\n", task1, task2), string(out))
	})
}
