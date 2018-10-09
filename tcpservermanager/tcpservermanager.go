package tcpservermanager

import (
	"fmt"
	"net"
	"sync"
)

var message string
var bufferSize int

type TcpServerManager struct {
	mtx sync.Mutex
}

func New() *TcpServerManager {
	return &TcpServerManager{}
}

func onServer() {
	bufferSize = 4096
	
	listener, err := net.Listen("tcp", ":2306")
	if err != nil {
		fmt.Printf("Error listen : %v\n", err)
		return
	}
	defer listener.Close()
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error Accept : %v\n")
			return
		}
		defer conn.Close()
		
		fmt.Printf("Connect remoteAddr : %v, localAddr : %v\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
		
		go onRead(conn)
		go onWrite(conn)
	}
}

func onRead(conn net.Conn) {
	data := make([]byte, bufferSize)
	for {
		n, err := conn.Read(data)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Printf("Discconect Conn : %v\n", err.Error())
				return
			}
			
			fmt.Printf("Error Read : %v\n", err)
			return
		}
		message = string(data[:n])
		fmt.Printf("Read : %v\n", message)
	}
}

func onWrite(conn net.Conn) {
	for {
		if len(message) <= 0 {
			continue
		}
		
		data := []byte(message)
		_, err := conn.Write(data)
		if err != nil {
			fmt.Printf("Error Write : %v\n", err)
			return
		}
		
		fmt.Printf("Write : %v\n", message)
		message = ""
	}
}

func (tm *TcpServerManager) RunTcpManager() {
	go onServer()
}