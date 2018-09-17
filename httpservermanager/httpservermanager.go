package httpservermanager

import (
	"net/http"
	"github.com/unrolled/render"
			"sync"
	"database/sql"
)

var renderer render.Render

type mysqlConnDBType int

const (
	MYSQL_Accountinfo mysqlConnDBType = iota
)

type httpManager struct {
	mtx sync.Mutex
	connMap map[mysqlConnDBType]*sql.DB
}

func New() *httpManager {
	return &httpManager{
		connMap:make(map[mysqlConnDBType]*sql.DB),
	}
}

func (hm *httpManager) makeMysqlConn(dbType mysqlConnDBType, conn *sql.DB) {
	if _, ok := hm.connMap[dbType]; !ok {
		hm.connMap[dbType] = conn
	}
}

func (hm *httpManager) getMysqlConn(dbType mysqlConnDBType) *sql.DB {
	if conn, ok := hm.connMap[dbType]; ok {
		return conn
	}
	return nil
}

func RunHttpServer() {
	hm := New()
	http.HandleFunc("/signup", hm.httpHandle_Signup)
	http.HandleFunc("/login", hm.httpHandle_Login)

	http.ListenAndServe(":2305", nil)
}