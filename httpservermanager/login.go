package httpservermanager

import (
	"database/sql"
	"fmt"
		"net/http"
	"encoding/json"
	"GoExample/gamedata"
)

func (hm *httpManager) httpHandle_Login(res http.ResponseWriter, req *http.Request) {
	data := make([]byte, 2048)
	n, _ := req.Body.Read(data)
	var req_pack req_LoginPacket
	json.Unmarshal([]byte(string(data[:n])), &req_pack)

	fmt.Printf("req = %v\n", string(data[:n]))

	res_pack := hm.call_Select_Login(req_pack)
	bytes, _ := json.Marshal(res_pack)

	renderer.Data(res, http.StatusOK, bytes)
}

func (hm *httpManager) call_Select_Login(req req_LoginPacket) rsp_LoginPacket {

	hm.mtx.Lock()
	defer hm.mtx.Unlock()

	rsp := rsp_LoginPacket{}

	conn, ok := hm.connMap[MYSQL_Accountinfo]
	if !ok {
		conn1, err := sql.Open("mysql", "root:ball2305@tcp(localhost:3306)/accountinfo")
		if err != nil {
			fmt.Printf("Error open mysql : %v\n", err)
			rsp.Error = 0
			return rsp
		}
		hm.makeMysqlConn(MYSQL_Accountinfo, conn1)
		conn = conn1
	}
	//defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		rsp.Error = gamedata.EC_MysqlConnectFail
		fmt.Println(err)
		return rsp
	}
	defer tx.Rollback()

	rows, err := conn.Query("select * from accountinfo where id=? && logintype=?", req.Id, req.LoginType)
	if err != nil {
		rsp.Error = gamedata.EC_UnknownError
		fmt.Println(err)
		return rsp
	}

	rowCnt := 0
	loginType := 0
	for rows.Next() {
		err := rows.Scan(&rsp.Uid, &rsp.Id, &rsp.Pw, &rsp.Nickname, &loginType)
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			fmt.Println(err)
			return rsp
		}
		rowCnt++
	}
	if rowCnt == 0 || rowCnt > 1 {
		rsp.Error = gamedata.EC_NotFoundAccount
		fmt.Println(err)
		return rsp
	}

	return rsp
}