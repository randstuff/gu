package guexec

import (
	"fmt"
	"testing"
)

func TestRunAndWriteOutput(t *testing.T) {

	tArgs := []string{"-lah", "/tmp"}
	out, err := RunAndWriteOutput("ls", tArgs, "/tmp/out")

	if err == nil {
		fmt.Println(out)
	}
}

func TestRun(t *testing.T) {

	tArgs := []string{"-lah", "/tmp"}
	out, err := Run("ls", tArgs)

	if err == nil {
		fmt.Println(out)
	}
}

func TestRunWithErr(t *testing.T) {

	tArgs := []string{"-lah", "/tmp"}
	out, strerr, err := RunWithErr("ls", tArgs)

	if err == nil {
		fmt.Println(strerr)
		fmt.Println(out)
	}
}
