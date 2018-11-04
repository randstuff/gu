package gulazy

import (
	"fmt"
	"testing"
)

func TestIntToString(t *testing.T) {

	fmt.Println(IntToString(42))
}

func TestFloat64ToString(t *testing.T) {
	fmt.Println(Float64ToString(42.42))
}

func TestStringToFloat64(t *testing.T) {

	fmt.Println(StringToFloat64("42.42"))
}

func TestTruncateFloat(t *testing.T) {

	fmt.Println(TruncateFloat(42.123456, 1))
	fmt.Println(TruncateFloat(42.123456, 2))
	fmt.Println(TruncateFloat(42.123456, 3))
	fmt.Println(TruncateFloat(42.123456, 4))
	fmt.Println(TruncateFloat(42.123456, 5))
	fmt.Println(TruncateFloat(42.123456, 6))
	fmt.Println(TruncateFloat(42.123456, 7))
}
