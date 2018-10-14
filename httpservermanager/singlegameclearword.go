package httpservermanager

import (
	"GoExample/gamedata"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func (hm *HttpServerManager) HttpHandle_SingleGameClearWord(res http.ResponseWriter, req *http.Request) {
	hm.mtx.Lock()
	defer hm.mtx.Unlock()
	
	data := make([]byte, 2048)
	n, _ := req.Body.Read(data)
	var req_pack Req_SingleGameClearWordPacket
	json.Unmarshal([]byte(string(data[:n])), &req_pack)
	
	fmt.Printf("singlegameclearword req = %v\n", string(data[:n]))
	
	res_pack, err := hm.call_Select_SingleGameClearWord(req_pack)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("singlegameclearword rsp : %v\n", res_pack)
	}
	bytes, _ := json.Marshal(res_pack)
	
	renderer.Data(res, http.StatusOK, bytes)
}

func (hm *HttpServerManager) call_Select_SingleGameClearWord(req Req_SingleGameClearWordPacket) (*Rsp_SingleGameClearWordPacket, error) {
	rsp := &Rsp_SingleGameClearWordPacket{}
	
	conn, ok := hm.connMap[MYSQL_SingleGameClearWord]
	if !ok {
		conn1, err := sql.Open("mysql", "root:ball2305@tcp(localhost:3306)/englishwordgame")
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			return rsp, fmt.Errorf("singlegameclearword error mysql conn : %v", err)
		}
		hm.makeMysqlConn(MYSQL_SingleGameClearWord, conn1)
		conn = conn1
	}
	
	tx, err := conn.Begin()
	if err != nil {
		rsp.Error = gamedata.EC_MysqlConnectFail
		return rsp, fmt.Errorf("singlegameclearword error mysql conn begin : %v", err)
	}
	defer tx.Rollback()
	
	rows, err := conn.Query("select english, clearcount, fastestcleartime from singlegameclearword where uid=? and english=?", req.Uid, req.English)
	if err != nil {
		rsp.Error = gamedata.EC_UnknownError
		return rsp, fmt.Errorf("singlegameclearword error mysql select : %v", err)
	}
	
	rowCnt := 0
	for rows.Next() {
		err := rows.Scan(&rsp.English, &rsp.ClearCount, &rsp.FastestClearTime)
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			fmt.Println(err)
			return rsp, fmt.Errorf("singlegameclearword error mysql scan : %v", err)
		}
		rowCnt++
	}
	
	if rowCnt == 1 {
		if req.FastestClearTime < rsp.FastestClearTime {
			rsp.FastestClearTime = req.FastestClearTime
		}
		result, err := conn.Exec("update singlegameclearword set clearcount=?, fastestcleartime=? where uid=? && english=?",
			rsp.ClearCount+1,
			rsp.FastestClearTime,
			req.Uid,
			req.English,
		)
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			return rsp, fmt.Errorf("singlegameclearword error mysql update : %v", err)
		}
		_ = result
	} else if rowCnt == 0 {
		rsp1, err := hm.call_Insert_SingleGameClearWord(req)
		if err != nil {
			return rsp1, err
		}
		rsp = rsp1
	} else {
		rsp.Error = gamedata.EC_NotFoundTableInfo
		return rsp, fmt.Errorf("singlegameclearword error %v", err)
	}
	
	return rsp, nil
}

func (hm *HttpServerManager) call_Insert_SingleGameClearWord(req Req_SingleGameClearWordPacket) (*Rsp_SingleGameClearWordPacket, error) {
	rsp := &Rsp_SingleGameClearWordPacket{}
	
	conn, ok := hm.connMap[MYSQL_SingleGameClearWord]
	if !ok {
		conn1, err := sql.Open("mysql", "root:ball2305@tcp(localhost:3306)/englishwordgame")
		if err != nil {
			rsp.Error = gamedata.EC_UnknownError
			return rsp, fmt.Errorf("singlegameclearword error mysql conn : %v", err)
		}
		hm.makeMysqlConn(MYSQL_SingleGameClearWord, conn1)
		conn = conn1
	}
	
	result, err := conn.Exec("insert into singlegameclearword (uid, english, clearcount, fastestcleartime) values (?, ?, ?, ?)",
		req.Uid,
		req.English,
		1,
		req.FastestClearTime,
	)
	if err != nil {
		rsp.Error = gamedata.EC_Table_Insert
		return rsp, fmt.Errorf("singlegameclearword error mysql insert : %v", err)
	}
	_ = result
	
	rsp.Error = gamedata.EC_Success
	rsp.English = req.English
	rsp.ClearCount = 1
	rsp.FastestClearTime = req.FastestClearTime
	
	return rsp, nil
}
