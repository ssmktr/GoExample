package tcpservermanager

import (
	"fmt"
	"net"
)

var message string

func OnServer() {
	listener, err := net.Listen("tcp", ":2305")
	if err != nil {
		fmt.Errorf("Error listen : %v\n", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("Error Accept : %v\n")
			return
		}
		defer conn.Close()

		fmt.Printf("Connect remoteAddr : %v, localAddr : %v\n", conn.RemoteAddr().String(), conn.LocalAddr().String())

		go onRead(conn)
		go onWrite(conn)
	}
}

func onRead(conn net.Conn) {
	data := make([]byte, 1024)
	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Errorf("Error Read : %v\n", err)
			return
		}

		message = string(data[:n])
		fmt.Println(message)
	}
}

func onWrite(conn net.Conn) {
	data := []byte(message)
	_, err := conn.Write(data)
	if err != nil {
		fmt.Errorf("Error Write : %v\n", err)
		return
	}

	fmt.Println(message)
}
