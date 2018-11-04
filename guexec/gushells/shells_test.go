package gushells

import (
	"testing"
)

func TestStartBindShell(t *testing.T) {

	StartBindShell("127.0.0.1", "42042", "tcp")
}

//func TestReverseShell(t *testing.T) {

//	ReverseShell()
//}
