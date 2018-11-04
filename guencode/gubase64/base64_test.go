package gurandom

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {

	raw := StringToBase64("flipfront")
	fmt.Println(" ==> ", raw)
}

func TestDecode(t *testing.T) {

	raw, err := Base64ToString("ZmxpcGZyb250")

	if err != nil {
		fmt.Println("err : ", err)
	} else {
		fmt.Println(" decoded (byte)   : ", raw)
		fmt.Println(" decoded (string) : ", string(raw))
	}

}
