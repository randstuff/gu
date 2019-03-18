package guhash

import (
	"fmt"
	"testing"
)

func TestMD5File(t *testing.T) {

	path := "/tmp/test"
	fmt.Println(MD5File(path))
}

func TestSHA1File(t *testing.T) {

	path := "/tmp/test"
	fmt.Println(SHA1File(path))
}

func TestSHA256File(t *testing.T) {

	path := "/tmp/test"
	fmt.Println(SHA256File(path))
}

func TestSHA512File(t *testing.T) {

	path := "/tmp/test"
	fmt.Println(SHA512File(path))
}
