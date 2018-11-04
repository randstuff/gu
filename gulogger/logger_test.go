package gulogger

import (
	"testing"
)

func TestInitLogger(t *testing.T) {

	InitLogger()

	Error.Println("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam sit amet aliquet sapien.  ")
	Warning.Println("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam sit amet aliquet sapien.  ")
	Info.Println("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam sit amet aliquet sapien.  ")
	Trace.Println("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam sit amet aliquet sapien.  ")

	EntryPoint.Println("LoremIpsum")
}
