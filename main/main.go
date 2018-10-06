package main

import (
	"GoExample/gametabledata"
	"GoExample/httpservermanager"
	_ "github.com/go-sql-driver/mysql"
	)

type gameManager struct {
	gameTableManager *gametabledata.GameTableDataManager
	httpServerManager *httpservermanager.HttpServerManager
}

func (gm *gameManager) New() {
	gm.gameTableManager = gametabledata.New()
	gm.httpServerManager = httpservermanager.New()
}

func main() {
	gm := &gameManager{}
	gm.New()
	
	gm.gameTableManager.RunGameTableDataServer(func() {
		// gm.httpServerManager.RunHttpServer()
	})
}

