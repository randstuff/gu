package gufilesop

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestWalk(t *testing.T) {

	tDir := NormalizePath("/tmp/")

	for w := range Walk(tDir) {
		for _, f := range w.Files {
			fp := filepath.Join(w.Dirpath, f.Name())
			fmt.Println(" :: ", fp, " :: ", filepath.Ext(fp))
		}
	}
}
