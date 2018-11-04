package guhash

import (
	"fmt"
	"testing"
)

const tString = "0xb4b3b4b3"

func TestMD5string(t *testing.T) {
	fmt.Println(MD5string(tString))
}

func TestSHA256string(t *testing.T) {
	fmt.Println(SHA256string(tString))
}

func TestSHA1string(t *testing.T) {
	fmt.Println(SHA1string(tString))
}

func TestMD5bytes(t *testing.T) {
	tByte := make([]byte, 512+1e6)
	fmt.Println(MD5bytes(tByte))
}

func TestSHA1bytes(t *testing.T) {
	tByte := make([]byte, 512+1e6)
	fmt.Println(SHA1bytes(tByte))
}

func TestSHA256bytes(t *testing.T) {
	tByte := make([]byte, 512+1e6)
	fmt.Println(SHA256bytes(tByte))
}
