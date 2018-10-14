package main

import (
	"GoExample/gametabledata"
	"GoExample/httpservermanager"
	"GoExample/tcpservermanager"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	gtm := gametabledata.New()
	
	gtm.RunGameTableDataServer(func() {
		hsm := httpservermanager.New()
		tsm := tcpservermanager.New()
		tsm.RunTcpManager()
		
		http.HandleFunc("/load_englishworddata", gtm.HttpHandle_load_englishworddata)
		http.HandleFunc("/load_localizationdata", gtm.HttpHandle_load_localizationdata)
		
		http.HandleFunc("/auth", hsm.HttpHandle_Auth)
		http.HandleFunc("/login", hsm.HttpHandle_Login)
		http.HandleFunc("/getuserinfo", hsm.HttpHandle_GetUserInfo)
		http.HandleFunc("/singlegameclearword", hsm.HttpHandle_SingleGameClearWord)
		
		fmt.Println("Start Http Server...")
		
		http.ListenAndServe(":2305", nil)
	})
}
