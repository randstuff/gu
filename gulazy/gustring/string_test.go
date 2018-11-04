package gustring

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {

	s := "Hello"
	fmt.Println(s)
	fmt.Println(Reverse(s))

}

func TestReverseUsingRunes(t *testing.T) {

	s := "Hello"
	fmt.Println(s)
	fmt.Println(ReverseUsingRunes(s))

}
