package httpservermanager

import (
	"database/sql"
	"fmt"
		"net/http"
	"encoding/json"
	"GoExample/gamedata"
	"time"
)

func (hm *httpManager) httpHandle_Login(res http.ResponseWriter, req *http.Request) {
	data := make([]byte, 2048)
	n, _ := req.Body.Read(data)
	var req_pack req_LoginPacket
	json.Unmarshal([]byte(string(data[:n])), &req_pack)

	fmt.Printf("req = %v\n", string(data[:n]))

	res_pack, err := hm.call_Select_Login(req_pack)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("rsp : %v\n", res_pack)
	}
	bytes, _ := json.Marshal(res_pack)

	renderer.Data(res, http.StatusOK, bytes)
}

func (hm *httpManager) call_Select_Login(req req_LoginPacket) (rsp_LoginPacket, error) {
	hm.mtx.Lock()
	defer hm.mtx.Unlock()

	rsp := rsp_LoginPacket{}

	conn, ok := hm.connMap[MYSQL_Accountinfo]
	if !ok {
		conn1, err := sql.Open("mysql", "root:ball2305@tcp(localhost:3306)/englishwordgame")
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			return rsp, fmt.Errorf("Error mysql conn : %v", err)
		}
		hm.makeMysqlConn(MYSQL_Accountinfo, conn1)
		conn = conn1
	}

	tx, err := conn.Begin()
	if err != nil {
		rsp.Error = gamedata.EC_MysqlConnectFail
		return rsp, fmt.Errorf("Error mysql conn begin : %v", err)
	}
	defer tx.Rollback()

	rows, err := conn.Query("select uid, id, lastlogindate from accountinfo where id=? && logintype=?", req.Id, req.LoginType)
	if err != nil {
		rsp.Error = gamedata.EC_UnknownError
		return rsp, fmt.Errorf("Error mysql select : %v", err)
	}

	curDate := time.Now().UTC()
	rowCnt := 0
	for rows.Next() {
		err := rows.Scan(&rsp.Uid, &rsp.Id, &rsp.Lastlogindate)
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			fmt.Println(err)
			return rsp, fmt.Errorf("Error mysql scan : %v", err)
		}
		rowCnt++
	}

	result, err := conn.Exec("update accountinfo set lastlogindate=? where uid=?", curDate, rsp.Uid)
	if err != nil {
		rsp.Error = gamedata.EC_UnknownError
		return rsp, fmt.Errorf("Error mysql update : %v", err)
	}
	_ = result
	rsp.Lastlogindate = curDate.String()

	if rowCnt == 0 || rowCnt > 1 {
		rsp.Error = gamedata.EC_NotFoundAccount
		return rsp, fmt.Errorf("Error %v", err)
	}

	return rsp, nil
}