package gufilesop

import (
	"fmt"
	"testing"
)

func TestMoveFile(t *testing.T) {

	src := "/tmp/a"
	dst := "/tmp/abc"

	if err := MoveFile(NormalizePath(src), NormalizePath(dst)); err != nil {
		fmt.Println(" err  : ", err)
	}
}
