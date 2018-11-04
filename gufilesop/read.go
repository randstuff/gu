package gufilesop

import (
	"bufio"
	"fmt"

	//	"log"
	"os"
)

func ReadFile(tPath string) {

	fd, err := os.Open(tPath)
	if err == nil {
		scanner := bufio.NewScanner(fd)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
	defer fd.Close()
}
