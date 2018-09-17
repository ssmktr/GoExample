package httpservermanager

import (
	"database/sql"
	"fmt"
		"net/http"
	"encoding/json"
	"GoExample/gamedata"
)

func (hm *httpManager) httpHandle_Signup(res http.ResponseWriter, req *http.Request) {
	data := make([]byte, 2048)
	n, _ := req.Body.Read(data)
	var req_pack req_SignupPacket
	json.Unmarshal([]byte(string(data[:n])), &req_pack)

	fmt.Printf("req = %v\n", string(data[:n]))

	res_pack := hm.call_Insert_Signup(req_pack)
	bytes, _ := json.Marshal(res_pack)

	renderer.Data(res, http.StatusOK, bytes)
}

func (hm *httpManager) call_Insert_Signup(req req_SignupPacket) rsp_SignupPacket {

	rsp := rsp_SignupPacket{}

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

	if rows.Next() {
		rsp.Error = gamedata.EC_AlreadyAccount
		fmt.Println(err)
		return rsp
	}

	hm.mtx.Lock()
	defer hm.mtx.Unlock()
	result, err := conn.Exec("insert into accountinfo (uid, id, pw, nickname, logintype) values (?, ?, ?, ?, ?)",
		req.Uid,
		req.Id,
		req.Pw,
		req.Nickname,
		gamedata.LT_LoginType(req.LoginType),
		)
	if err != nil {
		rsp.Error = gamedata.EC_UnknownError
		fmt.Println(err)
		return rsp
	}
	_ = result

	rsp.Error = gamedata.EC_Success
	rsp.Uid = req.Uid
	rsp.Id = req.Id
	rsp.Pw = req.Pw
	rsp.Nickname = req.Nickname

	return rsp
}