package gutime

import (
	"fmt"
	"testing"
)

func TestGetHourMinSec(t *testing.T) {
	fmt.Println(GetHourMinSec(true))
	fmt.Println(GetHourMinSec(false))
}

func TestGetHourMinSecAsInt(t *testing.T) {
	fmt.Println(GetHourMinSecAsInt())
}
