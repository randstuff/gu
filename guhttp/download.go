package guhttp

import (
	"fmt"
	"io"
	"net/http"

	//	"net/http"
	"gucrypto/hash"
	"os"
)

func DownloadFile(url string, FilePath string) string {

	defaultFileLocation := "N/A"
	SetProxy("")

	//      cDir := Pwd()
	cHashURL := guhash.MD5string(url)

	fileLocation := FilePath + cHashURL + ".jpg"

	// fmt.Println("Download", url, "to", fileName)

	output, err := os.Create(fileLocation)
	if err != nil {
		fmt.Println("Err :: File creation  ", fileLocation, "-", err)
		return defaultFileLocation
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Err :: Downloading", url, "-", err)
		return defaultFileLocation
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Err :: copying ", url, "-", err)
		return defaultFileLocation
	} else {
		return "/" + fileLocation
	}
	//      fmt.Println("FileLocation : ", fileLocation)
	//      fmt.Println("FileLocation : ", defaultFileLocation)

	return defaultFileLocation
}
