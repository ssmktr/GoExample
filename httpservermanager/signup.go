package httpservermanager

import (
	"database/sql"
	"fmt"
		"net/http"
	"encoding/json"
	"GoExample/gamedata"
	"time"
)

func (hm *httpManager) httpHandle_Signup(res http.ResponseWriter, req *http.Request) {
	data := make([]byte, 2048)
	n, _ := req.Body.Read(data)
	var req_pack req_SignupPacket
	json.Unmarshal([]byte(string(data[:n])), &req_pack)

	fmt.Printf("req = %v\n", string(data[:n]))

	res_pack, err := hm.call_Insert_Signup(req_pack)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("rsp : %v\n", res_pack)
	}
	bytes, _ := json.Marshal(res_pack)

	renderer.Data(res, http.StatusOK, bytes)
}

func (hm *httpManager) call_Insert_Signup(req req_SignupPacket) (rsp_SignupPacket, error) {
	hm.mtx.Lock()
	defer hm.mtx.Unlock()

	rsp := rsp_SignupPacket{}

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

	rows, err := conn.Query("select * from accountinfo where id=? && logintype=?", req.Id, req.LoginType)
	if err != nil {
		rsp.Error = gamedata.EC_UnknownError
		fmt.Printf("Error mysql select : %v", err)
		return rsp, fmt.Errorf("Error mysql select : %v", err)
	}

	if rows.Next() {
		rsp.Error = gamedata.EC_AlreadyAccount
		return rsp, fmt.Errorf("Error alreay account id : %v, logintype : %v", req.Id, req.LoginType)
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
		return rsp, fmt.Errorf("Error mysql insert : %v", err)
	}
	_ = result

	rsp.Error = gamedata.EC_Success
	rsp.Uid = req.Uid
	rsp.Id = req.Id
	rsp.Lastlogindate = curDate
	rsp.Createdate = curDate

	return rsp, nil
}