package gufilesop

import (
	"fmt"
	"testing"
)

func TestIsItExecutable(t *testing.T) {
	fmt.Println(IsItExecutable("ls"))
}

func TestPwd(t *testing.T) {

	if out, err := Pwd(); err == nil {
		fmt.Println(" pwd : ", out)
	}
}

func TestMkdir(t *testing.T) {

	if err := Mkdir(NormalizePath("/tmp/ThisIsATest")); err == nil {
		fmt.Println(" directory created.")
	} else {
		fmt.Println(" :: err - ", err)
	}
}

func TestListDirAndFiles(t *testing.T) {

	if files, dir, filesFull, dirFull, err := ListDirAndFiles(NormalizePath("/tmp/")); err == nil {

		fmt.Println(" files : ", files)
		fmt.Println(" dir   : ", dir)

		fmt.Println(" files FULL : ", filesFull)
		fmt.Println(" dir   FULL : ", dirFull)

	} else {
		fmt.Println(" err : ", err)
	}
}

func TestListFiles(t *testing.T) {

	if files, filesFull, err := ListFiles(NormalizePath("/tmp/")); err == nil {

		fmt.Println(" files : ", files)
		fmt.Println(" files FULL : ", filesFull)

	} else {
		fmt.Println(" err : ", err)
	}
}

func TestIsItFile(t *testing.T) {

	tPath := NormalizePath("/tmp/ThisIsATest/bbbb")

	if out := IsFile(tPath); out == true {
		fmt.Println(" is file : ", tPath)
	} else {
		fmt.Println(" something in the way ")
	}
}
