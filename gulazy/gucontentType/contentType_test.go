package guvarious

import (
	"fmt"
	"testing"
)

func TestGetFileContentType(t *testing.T) {

	out, err := GetFileContentType("/etc/passwd")
	if err != nil {
		fmt.Println(" err : ", err)
	} else {
		fmt.Println("out : ", out)
	}

}
