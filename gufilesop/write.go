package gufilesop

import (
	//	"fmt"
	"io/ioutil"
	"os"
)

func WritetoFile(tPath string, tData []byte) error {

	fd, err := os.Create(tPath)
	defer fd.Close()

	if err == nil {
		err = ioutil.WriteFile(tPath, tData, 0644)
	}
	return err
}
