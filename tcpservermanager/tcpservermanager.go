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
	
	connMap map[int][]net.Conn // [channel]
}

func New() *TcpServerManager {
	return &TcpServerManager{
		connMap: make(map[int][]net.Conn),
	}
}

func (tm *TcpServerManager) addConn(_channel int, _conn net.Conn) {
	if(len(tm.connMap[_channel]) >= 50) {
		fmt.Println("Error empty channel max count 50")
		return
	}
	
	tm.connMap[_channel] = append(tm.connMap[_channel], _conn)
	go tm.onRead(_conn)
	go tm.onWrite(_conn)
}

func (tm *TcpServerManager) onServer() {
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
		
		tm.addConn(1, conn)
	}
}

func (tm *TcpServerManager) onRead(conn net.Conn) {
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

func (tm *TcpServerManager) onWrite(conn net.Conn) {
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
	go tm.onServer()
}
