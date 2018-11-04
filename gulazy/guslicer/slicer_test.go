package guslices

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {

	target := []string{"aaa", "aaa", "bbb", "ccc"}
	fmt.Println(RemoveDuplicate(target))
}

func TestPrintSlice(t *testing.T) {

	target := []string{"aaa", "aaa", "bbb", "ccc"}
	PrintSlice(target)
}

func TestStringInSlice(t *testing.T) {
	target := []string{"aaa", "aaa", "bbb", "ccc"}

	if out := IsStringInSlice("bbb", target); out == true {
		fmt.Println(" 0 # Found ! ")
	} else {
		fmt.Println(" 0 #  NOT Found ! ")
	}

	if out := IsStringInSlice("bbzzzb", target); out == true {
		fmt.Println(" 1 # Found ! ")
	} else {
		fmt.Println(" 1 # NOT Found ! ")
	}

}
