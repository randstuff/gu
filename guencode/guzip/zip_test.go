package guzip

import (
	"testing"
)

func TestZip(t *testing.T) {
	Zip("/tmp/checks/", "/tmp/out.zip")
}

func TestUnzip(t *testing.T) {
	Unzip("/tmp/out.zip", "/tmp/out2/")
}
