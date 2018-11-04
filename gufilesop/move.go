package gufilesop

import (
	"fmt"
	"os"
)

func MoveFile(src string, dst string) error {

	if err := os.Rename(NormalizePath(src), NormalizePath(dst)); err != nil {
		fmt.Println(err)
		return nil
	} else {
		return err
	}

}
