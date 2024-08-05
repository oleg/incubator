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
	task := "test task number 1"

	dir, err := os.Getwd()
	assert.NoError(t, err)

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTaskFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		err := cmd.Run()
		assert.NoError(t, err)
	})
	task2 := "test task number 2"
	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		cmdStdIn, err := cmd.StdinPipe()
		assert.NoError(t, err)
		io.WriteString(cmdStdIn, task2)
		cmdStdIn.Close()
		err = cmd.Run()
		assert.NoError(t, err)
	})
	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("  1: %s\n\n", task), string(out))
	})
}
