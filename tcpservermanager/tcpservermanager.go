package tcpservermanager

import (
	"fmt"
	"net"
	"sync"
	"time"
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
	
	if (len(tm.connMap[_channel]) >= 50) {
		fmt.Println("Error empty channel max count 50")
		return
	}
	
	if _, ok := tm.connMap[_channel][_conn]; ok {
		return
	}
	
	tm.connMap[_channel][_conn] = true
	
	tm.onRead(_conn)
	tm.onWrite(_conn)
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

func (tm *TcpServerManager) getUserCountByChannel(_channel int) int {
	if conns, ok := tm.connMap[_channel]; ok {
		return len(conns)
	}
	
	return 0
}

func (tm *TcpServerManager) onRead(conn net.Conn) {
	data := make([]byte, bufferSize)
	fmt.Println(tm.getUserCountByChannel(1))
	for {
		n, err := conn.Read(data)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Printf("Discconect Conn : %v\n", err.Error())
				tm.leaveConn(conn)
				fmt.Println(tm.getUserCountByChannel(1))
				return
			}
			
			fmt.Printf("Error Read : %v\n", err)
			return
		}
		message = string(data[:n])
		fmt.Printf("Read : %v\n", message)
		
		time.Sleep(time.Second * 1)
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
		
		time.Sleep(time.Second * 1)
	}
}

func (tm *TcpServerManager) RunTcpManager() {
	go tm.onServer()
}
