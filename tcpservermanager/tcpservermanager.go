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
	
	connMap map[int]map[net.Conn]bool // [channel][conn]
}

func New() *TcpServerManager {
	return &TcpServerManager{
		connMap: make(map[int]map[net.Conn]bool),
	}
}

func (tm *TcpServerManager) addConn(_channel int, _conn net.Conn) {
	if _, ok := tm.connMap[_channel]; !ok {
		tm.connMap[_channel] = make(map[net.Conn]bool)
	}
	
	fmt.Println("===========11")
	
	if (len(tm.connMap[_channel]) >= 50) {
		fmt.Println("Error empty channel max count 50")
		return
	}
	
	fmt.Println("===========22")
	
	if _, ok := tm.connMap[_channel][_conn]; ok {
		return
	}
	
	fmt.Println("===========33")
	
	tm.connMap[_channel][_conn] = true
	go tm.onRead(_conn)
	go tm.onWrite(_conn)
}

func (tm *TcpServerManager) leaveConn(_conn net.Conn) {
	for cha, conns := range tm.connMap {
		if _, ok := tm.connMap[cha][_conn]; ok {
			delete(conns, _conn)
			break
		}
	}
}

func (tm *TcpServerManager) onServer() {
	bufferSize = 4096
	
	listener, err := net.Listen("tcp", ":2306")
	if err != nil {
		fmt.Printf("Error listen : %v\n", err)
		return
	}
	defer listener.Close()
	
	fmt.Println("Start Tcp Server...")
	
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
	
	fmt.Println("==========44")
	
	data := make([]byte, bufferSize)
	fmt.Println(len(tm.connMap[1]))
	for {
		
		fmt.Println("==========55")
		
		n, err := conn.Read(data)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Printf("Discconect Conn : %v\n", err.Error())
				tm.leaveConn(conn)
				fmt.Println(len(tm.connMap[1]))
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
