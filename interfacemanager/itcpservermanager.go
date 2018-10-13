package interfacemanager

import (
	"GoExample/tcpservermanager"
)

type ITcpServerManager interface {
	LeaveConn(_user *tcpservermanager.User)
}
