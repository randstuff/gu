package guexec

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func RunAndWriteOutput(tBin string, tArg []string, tOutput string) ([]byte, error) {

	out, err := exec.Command(tBin, tArg...).Output()
	if err != nil {
		log.Fatal(err)
	} else {
		_ = WritetoFile(tOutput, out)
	}

	return out, err
}

func Run(tBin string, tArg []string) ([]byte, error) {

	out, err := exec.Command(tBin, tArg...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out, err
}

func RunWithErr(tBin string, tArg []string) ([]byte, []byte, error) {

	var stdout, stderr bytes.Buffer

	cmd := exec.Command(tBin, tArg...)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	out, tStdErr := stdout.Bytes(), stderr.Bytes()

	return out, tStdErr, err
}

// //

func WritetoFile(tPath string, tData []byte) error {

	fd, err := os.Create(tPath)
	defer fd.Close()

	if err == nil {
		err = ioutil.WriteFile(tPath, tData, 0644)
	}
	return err
}
