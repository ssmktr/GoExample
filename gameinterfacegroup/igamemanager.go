package gameinterfacegroup

import (
	"GoExample/gametabledata"
	"GoExample/httpservermanager"
	"GoExample/tcpservermanager"
)

type IGameManager interface {
	GetGameTableManager() *gametabledata.GameTableDataManager
	GetHttpServerManager() *httpservermanager.HttpServerManager
	GetTcpServerManager() *tcpservermanager.TcpServerManager
}
