package main

import (
	"GoExample/gameinterfacegroup"
	"GoExample/gamemanager"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	gm := &gamemanager.GameManager{}
	gm.New()
	igm := gameinterfacegroup.IGameManager(gm)
	
	igm.GetGameTableManager().RunGameTableDataServer(func() {
		gm.HttpServerManager.RunHttpServer()
		gm.TcpServerManager.RunTcpManager()
		
		http.HandleFunc("/load_englishworddata", gm.GameTableManager.HttpHandle_load_englishworddata)
		http.HandleFunc("/load_localizationdata", gm.GameTableManager.HttpHandle_load_localizationdata)
		
		http.HandleFunc("/auth", gm.HttpServerManager.HttpHandle_Auth)
		http.HandleFunc("/login", gm.HttpServerManager.HttpHandle_Login)
		http.HandleFunc("/getuserinfo", gm.HttpServerManager.HttpHandle_GetUserInfo)
		
		fmt.Println("Start Http Server...")
		
		http.ListenAndServe(":2305", nil)
	})
}
