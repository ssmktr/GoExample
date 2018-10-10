package gamemanager

import (
	"GoExample/gameinterfacegroup"
	"GoExample/gametabledata"
	"GoExample/httpservermanager"
	"GoExample/tcpservermanager"
)

var _ gameinterfacegroup.IGameManager = &GameManager{}

type GameManager struct {
	GameTableManager  *gametabledata.GameTableDataManager
	HttpServerManager *httpservermanager.HttpServerManager
	TcpServerManager  *tcpservermanager.TcpServerManager
}

func (gm *GameManager) New() {
	gm.GameTableManager = gametabledata.New()
	gm.HttpServerManager = httpservermanager.New()
	gm.TcpServerManager = tcpservermanager.New()
}
