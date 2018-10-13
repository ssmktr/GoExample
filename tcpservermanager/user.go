package tcpservermanager

import (
	"GoExample/gamedata"
	"fmt"
	"net"
	"sync"
)

type User struct {
	mtx     sync.Mutex
	conn    net.Conn
	channel int
	message string
	
	tcpServerManager *TcpServerManager
}

func NewUser(_conn net.Conn, _channel int, _tcpServerManager *TcpServerManager) *User {
	return &User{
		conn:             _conn,
		channel:          _channel,
		tcpServerManager: _tcpServerManager,
	}
}

func (u *User) Initialize() {
	go u.onRead()
	go u.onWrite()
}

func (u *User) onRead() {
	data := make([]byte, gamedata.BufferSize)
	for {
		n, err := u.conn.Read(data)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Printf("Discconect Conn : %v\n", err.Error())
				u.tcpServerManager.LeaveConn(u)
				return
			}
			
			fmt.Printf("Error Read : %v\n", err)
			return
		}
		u.message = string(data[:n])
		fmt.Printf("Read : %v\n", u.message)
	}
}

func (u *User) onWrite() {
	for {
		if len(u.message) <= 0 {
			continue
		}
		
		u.Send(u.message)
		
		data := []byte(u.message)
		_, err := u.conn.Write(data)
		if err != nil {
			fmt.Printf("Error Write : %v\n", err)
			continue
		}
		
		fmt.Printf("Write : %v\n", u.message)
		u.message = ""
	}
}

func (u *User) Send(_message string) {
	message := _message
	for _cha, _users := range u.tcpServerManager.ConnMap {
		if _cha == u.channel {
			for _user, _ := range _users {
				if u != _user {
					_user.Receive(message)
				}
			}
		}
	}
}

func (u *User) Receive(_message string) {
	u.message = _message
}
