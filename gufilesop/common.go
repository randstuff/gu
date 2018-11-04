package gufilesop

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func Pwd() (string, error) {

	if r, err := os.Getwd(); err == nil {
		fmt.Println(" :: pwd - ", r)
		return r, nil
	} else {
		return "N/A", err

	}
}

func Mkdir(path string) error {

	if _, err := os.Stat(NormalizePath(path)); os.IsNotExist(err) {
		os.Mkdir(path, 0744)
		return nil
	} else {
		return err
	}
}

func ListDirAndFiles(path string) ([]string, []string, []string, []string, error) {

	var lFiles []string
	var lDir []string

	var lFilesFull []string
	var lDirFull []string

	dirPath := NormalizePath(path)
	r, err := ioutil.ReadDir(NormalizePath(path))

	for _, f := range r {
		//		fmt.Println("=> ", f.Name())

		switch {
		case f.Mode().IsDir():
			lDir = append(lDir, f.Name())
			lDirFull = append(lDirFull, filepath.Join(dirPath, f.Name()))
		case f.Mode().IsRegular():
			lFiles = append(lFiles, f.Name())
			lFilesFull = append(lFilesFull, filepath.Join(dirPath, f.Name()))

			//	case f.Mode()&os.ModeSymlink == os.ModeSymlink :
		}

	}
	return lFiles, lDir, lFilesFull, lDirFull, err
}

func ListFiles(tPath string) ([]string, []string, error) {

	//	_ = filepath.Walk(tPath, func(path string, f os.FileInfo, err error) error {

	var lFiles []string
	var lFilesFull []string

	dirPath := NormalizePath(tPath)
	r, err := ioutil.ReadDir(NormalizePath(tPath))

	for _, f := range r {
		// fmt.Println("=> ", f.Name())

		if f.Mode().IsRegular() == true {
			lFiles = append(lFiles, f.Name())
			lFilesFull = append(lFilesFull, filepath.Join(dirPath, f.Name()))
		}
	}
	return lFiles, lFilesFull, err

}

func IsItExecutable(tPath string) (string, error) {

	r, err := exec.LookPath(NormalizePath(tPath))
	if err != nil {
		fmt.Println("not found : [", tPath, "]")
	} else {
		fmt.Println(" :: Found - ", r)
	}

	return r, err
}

func IsFile(tFile string) bool {

	stat, err := os.Stat(NormalizePath(tFile))
	if err == nil {
		if !stat.IsDir() {
			fmt.Println("ok 3")
			return true
		} else {
			return false
		}
	}
	return false
}
