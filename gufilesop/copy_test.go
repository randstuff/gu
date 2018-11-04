package gufilesop

import (
	"fmt"
	"testing"
)

func TestCopyDir(t *testing.T) {

	fmt.Println(".")

	if err := CopyDir(NormalizePath("/tmp/a/"), NormalizePath("/tmp/b/")); err != nil {
		fmt.Println(" err : ", err)
	}
}
