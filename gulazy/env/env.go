package guvarious

import (
	"fmt"
	"os"
	"strings"
)

func ListEnvVar() {

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], "=", pair[1])
	}
}
