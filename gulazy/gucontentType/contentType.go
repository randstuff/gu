package guvarious

import (
	"fmt"
	"net/http"
	"os"
)

func GetFileContentType(tFilePath string) (string, error) {

	f, err := os.Open(tFilePath)

	if err != nil {
		fmt.Println(" err : ", err)
	}
	defer f.Close()

	//	contentType, err := GetFileContentType(f)
	//	if err != nil {
	//		fmt.Println(" err : ", err)

	//	} else {

	buffer := make([]byte, 512)

	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	//}

	return contentType, nil
}
