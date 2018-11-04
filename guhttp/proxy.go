package guhttp

import (
	//	"net/http"
	"fmt"
	"os"
)

func SetProxy(tProxy string) {

	if tProxy != "" {
		os.Setenv("HTTP_PROXY", tProxy)
		os.Setenv("HTTPS_PROXY", tProxy)
		// os.Setenv("HTTP_PROXY", "127.0.0.1:3128")
		// os.Setenv("HTTPS_PROXY", "127.0.0.1:3128")
	}

	fmt.Println(" HTTP_PROXY  :", os.Getenv("HTTP_PROXY"))
	fmt.Println(" HTTPS_PROXY :", os.Getenv("HTTPS_PROXY"))

}
