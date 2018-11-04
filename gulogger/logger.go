package gulogger

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	Error   *log.Logger
	Warning *log.Logger
	Info    *log.Logger

	//

	Trace      *log.Logger
	EntryPoint *log.Logger

	//

	Clean *log.Logger
)

//var debug = 110110
var debug = 111111

func SetLogSettingstings(errorHandle io.Writer, warningHandle io.Writer, infoHandle io.Writer, traceHandle io.Writer, entryPointHandle io.Writer, CleanHandle io.Writer) {

	Error = log.New(errorHandle, color.HiRedString("#        ERROR : "), log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, color.YellowString("#      WARNING : "), log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, color.HiBlueString("#         INFO : "), log.Ldate|log.Ltime|log.Lshortfile)

	Trace = log.New(traceHandle, color.HiMagentaString("#        TRACE : "), log.Ldate|log.Ltime|log.Lshortfile)

	EntryPoint = log.New(entryPointHandle, color.CyanString(":: ENTRY POINT : "), log.Ldate|log.Ltime|log.Lshortfile)
}

func InitLogger() {

	switch debug {

	case 100000: // Error
		SetLogSettingstings(os.Stderr, ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)

	case 110000: // Error / Warning
		SetLogSettingstings(os.Stderr, os.Stdout, ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)

	case 111000: // Error / Warning / Info
		SetLogSettingstings(os.Stderr, ioutil.Discard, ioutil.Discard, os.Stdout, os.Stdout, os.Stdout)

	case 111100: // Error / Warning / Info / Trace
		SetLogSettingstings(os.Stderr, os.Stdout, os.Stdout, os.Stdout, ioutil.Discard, ioutil.Discard)

	case 111110: // Error / Warning / Info / Trace / EntryPoint
		SetLogSettingstings(os.Stderr, os.Stdout, os.Stdout, os.Stdout, os.Stdout, ioutil.Discard)

	case 110010: // Error / Warning / EntryPoint
		SetLogSettingstings(os.Stderr, os.Stdout, ioutil.Discard, ioutil.Discard, os.Stdout, ioutil.Discard)

	case 110110: // Error / Warning / EntryPoint / Trace
		SetLogSettingstings(os.Stderr, os.Stdout, ioutil.Discard, os.Stdout, os.Stdout, ioutil.Discard)

	case 111111: // Full
		SetLogSettingstings(os.Stderr, os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stdout)

	default:
		SetLogSettingstings(os.Stderr, ioutil.Discard, ioutil.Discard, os.Stdout, os.Stdout, os.Stdout)
	}
}

func Clean1Formated(tHrd string, tMsg string) {

	row := fmt.Sprintf(":: %-32s - [%-128s] - ", tHrd, tMsg)
	Clean.Println(row)
}
