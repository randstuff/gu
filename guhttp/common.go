package guhttp

import (
	//	"crypto/md5"
	//	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"

	"guregex"
	//	"regexp"
	//	"strconv"
)

func HttpGET(url string, cRegex string, tProxy string) ([]string, string) {

	SetProxy(tProxy)

	out := make([]string, 0)
	raw := ""

	if url != "N/A" && url != "" {

		res, err := http.Get(url)
		if err != nil {
			fmt.Println(":: Err - GET \"" + url + "\"")

		} else {

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(":: Err - ioutil.ReadAll", err)

			} else {
				//fmt.Println("HERE berfore strdata")
				strData := string(body)

				//fmt.Println("::", strData)
				out = guregex.RegFind(strData, cRegex)
				return out, strData
			}
		}
	}

	return out, raw
}
