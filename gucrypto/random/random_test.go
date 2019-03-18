package gurandom

import (
	"fmt"
	"testing"
)

func TestRandMD5(t *testing.T) {
	fmt.Println(RandMD5())
}

func TestRandSHA1(t *testing.T) {
	fmt.Println(RandSHA1())
}

func TestRandSHA256(t *testing.T) {
	fmt.Println(RandSHA256())
}

func TestRandSHA512(t *testing.T) {
	fmt.Println(RandSHA512())
}

func TestRandStringBytes(t *testing.T) {
	fmt.Println(RandStringBytes(42))
}

func TestRandBytes(t *testing.T) {
	fmt.Println(RandBytes(42))
}
