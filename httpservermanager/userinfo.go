package httpservermanager

import (
	"GoExample/gamedata"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func (hm *HttpServerManager) HttpHandle_GetUserInfo(res http.ResponseWriter, req *http.Request) {
	hm.mtx.Lock()
	defer hm.mtx.Unlock()
	
	data := make([]byte, 2048)
	n, _ := req.Body.Read(data)
	var req_pack Req_GetUserInfoPacket
	json.Unmarshal([]byte(string(data[:n])), &req_pack)
	
	fmt.Printf("userinfo req = %v\n", string(data[:n]))
	
	res_pack, err := hm.call_Select_UserInfo(req_pack)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("userinfo rsp : %v\n", res_pack)
	}
	bytes, _ := json.Marshal(res_pack)
	
	renderer.Data(res, http.StatusOK, bytes)
}

func (hm *HttpServerManager) call_Select_UserInfo(req Req_GetUserInfoPacket) (*Rsp_GetUserInfoPacket, error) {
	rsp := &Rsp_GetUserInfoPacket{}
	
	conn, ok := hm.connMap[MYSQL_UserInfo]
	if !ok {
		conn1, err := sql.Open("mysql", "root:ball2305@tcp(localhost:3306)/englishwordgame")
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			return rsp, fmt.Errorf("userinfo error mysql conn : %v", err)
		}
		hm.makeMysqlConn(MYSQL_UserInfo, conn1)
		conn = conn1
	}
	
	tx, err := conn.Begin()
	if err != nil {
		rsp.Error = gamedata.EC_MysqlConnectFail
		return rsp, fmt.Errorf("userinfo error mysql conn begin : %v", err)
	}
	defer tx.Rollback()
	
	rows, err := conn.Query("select nickname, energy, gold, heart from userinfo where uid=?", req.Uid)
	if err != nil {
		rsp.Error = gamedata.EC_UnknownError
		return rsp, fmt.Errorf("userinfo error mysql select : %v", err)
	}
	
	rowCnt := 0
	for rows.Next() {
		err := rows.Scan(&rsp.NickName, &rsp.Energy, &rsp.Gold, &rsp.Heart)
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			fmt.Println(err)
			return rsp, fmt.Errorf("userinfo error mysql scan : %v", err)
		}
		rowCnt++
	}
	
	if rowCnt == 0 || rowCnt > 1 {
		rsp.Error = gamedata.EC_NotFoundTableInfo
		return rsp, fmt.Errorf("userinfo error %v", err)
	}
	
	return rsp, nil
}

func (hm *HttpServerManager) call_Insert_UserInfo(uid, nickname string) error {
	conn, ok := hm.connMap[MYSQL_UserInfo]
	if !ok {
		conn1, err := sql.Open("mysql", "root:ball2305@tcp(localhost:3306)/englishwordgame")
		if err != nil {
			return fmt.Errorf("userinfo error mysql conn : %v", err)
		}
		hm.makeMysqlConn(MYSQL_UserInfo, conn1)
		conn = conn1
	}
	
	tx, err := conn.Begin()
	if err != nil {
		return fmt.Errorf("userinfo error mysql conn begin : %v", err)
	}
	defer tx.Rollback()
	
	result, err := conn.Exec("insert into userinfo (uid, nickname, energy, gold, heart) values (?, ?, ?, ?, ?)",
		uid,
		nickname,
		10,
		1000,
		0,
	)
	if err != nil {
		return fmt.Errorf("userinfo error mysql insert : %v", err)
	}
	_ = result
	
	return nil
}
