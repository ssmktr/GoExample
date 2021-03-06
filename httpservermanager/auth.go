package httpservermanager

import (
	"database/sql"
	"fmt"
		"net/http"
	"encoding/json"
	"GoExample/gamedata"
	"time"
)

func (hm *HttpServerManager) HttpHandle_Auth(res http.ResponseWriter, req *http.Request) {
	hm.mtx.Lock()
	defer hm.mtx.Unlock()
	
	data := make([]byte, 2048)
	n, _ := req.Body.Read(data)
	var req_pack Req_AuthPacket
	json.Unmarshal([]byte(string(data[:n])), &req_pack)

	fmt.Printf("auth req = %v\n", string(data[:n]))

	res_pack, err := hm.call_Select_Auth(req_pack)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("auth rsp : %v\n", res_pack)
	}
	bytes, _ := json.Marshal(res_pack)

	renderer.Data(res, http.StatusOK, bytes)
}

func (hm *HttpServerManager) call_Select_Auth(req Req_AuthPacket) (*Rsp_AuthPacket, error) {
	rsp := &Rsp_AuthPacket{}

	conn, ok := hm.connMap[MYSQL_Accountinfo]
	if !ok {
		conn1, err := sql.Open("mysql", "root:ball2305@tcp(localhost:3306)/englishwordgame")
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			return rsp, fmt.Errorf("auth error mysql conn : %v", err)
		}
		hm.makeMysqlConn(MYSQL_Accountinfo, conn1)
		conn = conn1
	}

	tx, err := conn.Begin()
	if err != nil {
		rsp.Error = gamedata.EC_MysqlConnectFail
		return rsp, fmt.Errorf("auth error mysql conn begin : %v", err)
	}
	defer tx.Rollback()

	rows, err := conn.Query("select uid, id, logintype, lastlogindate from accountinfo where id=? && logintype=?", req.Id, req.LoginType)
	if err != nil {
		rsp.Error = gamedata.EC_UnknownError
		fmt.Printf("Error mysql select : %v", err)
		return rsp, fmt.Errorf("auth error mysql select : %v", err)
	}

	rowsCount := 0
	for rows.Next() {
		err := rows.Scan(&rsp.Uid, &rsp.Id, &rsp.LoginType, &rsp.Lastlogindate)
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			fmt.Println(err)
			return rsp, fmt.Errorf("auth error mysql scan : %v", err)
		}
		rowsCount++
	}

	if rowsCount == 1 {
		curDate := time.Now().UTC()
		result, err := conn.Exec("update accountinfo set lastlogindate=? where id=? && logintype=?", curDate, req.Id, req.LoginType)
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			return rsp, fmt.Errorf("auth error mysql update : %v", err)
		}
		_ = result
		rsp.Lastlogindate = curDate.String()
	} else if rowsCount == 0 {
		rsp.Error = gamedata.EC_NotFoundAccount
		if rsp1, err := hm.call_Insert_Auth(req); err != nil {
			return rsp, err
		} else {
			rsp = rsp1
		}
	}

	return rsp, nil
}

func (hm *HttpServerManager) call_Insert_Auth(req Req_AuthPacket) (*Rsp_AuthPacket, error) {
	rsp := &Rsp_AuthPacket{}

	conn, ok := hm.connMap[MYSQL_Accountinfo]
	if !ok {
		conn1, err := sql.Open("mysql", "root:ball2305@tcp(localhost:3306)/englishwordgame")
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			return rsp, fmt.Errorf("auth error mysql conn : %v", err)
		}
		hm.makeMysqlConn(MYSQL_Accountinfo, conn1)
		conn = conn1
	}

	tx, err := conn.Begin()
	if err != nil {
		rsp.Error = gamedata.EC_MysqlConnectFail
		return rsp, fmt.Errorf("auth error mysql conn begin : %v", err)
	}
	defer tx.Rollback()
	
	if req.Uid == "" || req.Id == "" || req.NickName == "" || req.LoginType == gamedata.LT_None {
		rsp.Error = gamedata.EC_Table_Insert
		return rsp, fmt.Errorf("auth insert account uid : %v, id : %v, nickname : %v, logintype : %v", req.Uid, req.Id, req.NickName, req.LoginType)
	}

	curDate := time.Now().UTC()
	result, err := conn.Exec("insert into accountinfo (uid, id, logintype, lastlogindate, createdate) values (?, ?, ?, ?, ?)",
		req.Uid,
		req.Id,
		gamedata.LT_LoginType(req.LoginType),
		curDate,
		curDate,
		)
	if err != nil {
		rsp.Error = gamedata.EC_UnknownError
		return rsp, fmt.Errorf("auth error mysql insert : %v", err)
	}
	_ = result
	
	if err := hm.call_Insert_UserInfo(req.Uid, req.NickName); err != nil {
		rsp.Error = gamedata.EC_Table_Insert
		return rsp, err
	}

	rsp.Error = gamedata.EC_Success
	rsp.Uid = req.Uid
	rsp.Id = req.Id
	rsp.LoginType = req.LoginType
	rsp.Lastlogindate = curDate.String()

	return rsp, nil
}