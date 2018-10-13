package tcpservermanager

import (
	"fmt"
	"net"
	"sync"
)

type TcpServerManager struct {
	mtx sync.Mutex
	
	ConnMap map[int]map[*User]bool // [channel][User]
}

func New() *TcpServerManager {
	return &TcpServerManager{
		ConnMap: make(map[int]map[*User]bool),
	}
}

func (tm *TcpServerManager) addConn(_channel int, _user *User) {
	fmt.Printf("1. connected user count : %v\n", tm.getUserCountByChannel(1))
	
	if _, ok := tm.ConnMap[_channel]; !ok {
		tm.ConnMap[_channel] = make(map[*User]bool)
	}
	
	if (len(tm.ConnMap[_channel]) >= 50) {
		fmt.Println("Error empty channel max count 50")
		return
	}
	
	if _, ok := tm.ConnMap[_channel][_user]; ok {
		return
	}
	
	tm.ConnMap[_channel][_user] = true
	
	fmt.Printf("2. connected user count : %v\n", tm.getUserCountByChannel(1))
}

func (tm *TcpServerManager) LeaveConn(_user *User) {
	for cha, conns := range tm.ConnMap {
		if _, ok := tm.ConnMap[cha][_user]; ok {
			delete(conns, _user)
			break
		}
	}
	
	fmt.Printf("3. connected user count : %v\n", tm.getUserCountByChannel(1))
}

func (tm *TcpServerManager) getUserCountByChannel(_channel int) int {
	if conns, ok := tm.ConnMap[_channel]; ok {
		return len(conns)
	}
	
	return 0
}

func (tm *TcpServerManager) onServer() {
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
		
		user := NewUser(conn, 1, tm)
		user.Initialize()
		tm.addConn(1, user)
	}
}

func (tm *TcpServerManager) RunTcpManager() {
	go tm.onServer()
}
