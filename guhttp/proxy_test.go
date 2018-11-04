package guhttp

import (
	"testing"
)

func MoveFile(src string, dst string) bool {

	if err := os.Rename(NormalizePath(src), NormalizePath(dst)); err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}
