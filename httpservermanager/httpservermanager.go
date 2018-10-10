package httpservermanager

import (
	"GoExample/gameinterfacegroup"
	"database/sql"
	"github.com/unrolled/render"
	"sync"
)

var renderer render.Render

type mysqlConnDBType int

const (
	MYSQL_Accountinfo mysqlConnDBType = iota
	MYSQL_UserInfo
)

var _ gameinterfacegroup.IHttpServerManager = &HttpServerManager{}

type HttpServerManager struct {
	mtx     sync.Mutex
	connMap map[mysqlConnDBType]*sql.DB
}

func New() *HttpServerManager {
	return &HttpServerManager{
		connMap: make(map[mysqlConnDBType]*sql.DB),
	}
}

func (hm *HttpServerManager) makeMysqlConn(dbType mysqlConnDBType, conn *sql.DB) {
	if _, ok := hm.connMap[dbType]; !ok {
		hm.connMap[dbType] = conn
	}
}

func (hm *HttpServerManager) getMysqlConn(dbType mysqlConnDBType) *sql.DB {
	if conn, ok := hm.connMap[dbType]; ok {
		return conn
	}
	return nil
}

func (hm *HttpServerManager) RunHttpServer() {
	// http.HandleFunc("/auth", hm.HttpHandle_Auth)
	// http.HandleFunc("/login", hm.HttpHandle_Login)
	// http.HandleFunc("/getuserinfo", hm.HttpHandle_GetUserInfo)
	//
	// http.ListenAndServe(":2305", nil)
}
