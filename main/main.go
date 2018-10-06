package main

import (
	"GoExample/gametabledata"
	"GoExample/httpservermanager"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
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
		gm.httpServerManager.RunHttpServer()
		
		http.HandleFunc("/load_englishworddata", gm.gameTableManager.HttpHandle_load_englishworddata)
		
		http.HandleFunc("/auth", gm.httpServerManager.HttpHandle_Auth)
		http.HandleFunc("/login", gm.httpServerManager.HttpHandle_Login)
		http.HandleFunc("/getuserinfo", gm.httpServerManager.HttpHandle_GetUserInfo)
		
		http.ListenAndServe(":2305", nil)
	})
}

