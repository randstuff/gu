package gulazy

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

func IntToString(target int64) string {
	return strconv.FormatInt(target, 10) // , 64)
	//      (target, 'f', 6, 64)
}

func UInt64toString(target uint64) string {
	return strconv.FormatUint(target, 10)
}

func Float64ToString(target float64) string {
	return strconv.FormatFloat(target, 'f', 6, 64)
}

func StringToFloat64(mystring string) float64 {
	out, _ := strconv.ParseFloat(mystring, 64)
	return out
}

func TruncateFloat(some float64, floatCount int) float64 {

	switch floatCount {

	case 1:
		return float64(int(some*10)) / 10
	case 2:
		return float64(int(some*100)) / 100
	case 3:
		return float64(int(some*1000)) / 1000
	case 4:
		return float64(int(some*10000)) / 10000
	case 5:
		return float64(int(some*100000)) / 100000
	case 6:
		return float64(int(some*1000000)) / 1000000
	}
	return 0
}

func ReadCloserToString(t io.ReadCloser) string {

	buf := new(bytes.Buffer)
	buf.ReadFrom(t)
	newStr := buf.String()

	fmt.Printf(newStr)
	return newStr

}
