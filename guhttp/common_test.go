package guhttp

import (
	"testing"
)

func TestHttpGET(src string, dst string) bool {

	if err := os.Rename(NormalizePath(src), NormalizePath(dst)); err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}
