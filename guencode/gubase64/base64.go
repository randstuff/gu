package gurandom

import (
	"encoding/base64"
	"fmt"
)

func StringToBase64(t string) string {

	encoded := base64.StdEncoding.EncodeToString([]byte(t))
	return encoded
}

func Base64ToString(t string) ([]byte, error) {

	raw, err := base64.StdEncoding.DecodeString(t)

	if err != nil {
		fmt.Println("err : ", err)
	}
	return raw, err

}
