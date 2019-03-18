package guslices

import (
	"fmt"
	//	"testing"
)

func RemoveDuplicate(target []string) []string {

	cleaned := []string{}

	for _, value := range target {

		if !IsStringInSlice(value, cleaned) {
			cleaned = append(cleaned, value)
		}
	}
	return cleaned
}

func PrintSlice(slice []string) {

	for i := range slice {
		fmt.Println(i, slice[i])
	}
}

func IsStringInSlice(str string, list []string) bool {

	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
