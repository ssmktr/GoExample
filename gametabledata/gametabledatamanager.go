package gametabledata

import (
	"GoExample/gamedata"
	"GoExample/httpservermanager"
	"encoding/json"
	"fmt"
	"github.com/unrolled/render"
	"net/http"
	"sync"
)

var renderer render.Render

type GameTableDataManager struct {
	mtx sync.Mutex
	
	englishWordSlice    []*englishWordData
	localizationDataMap map[int]*localizationData
}

func New() *GameTableDataManager {
	return &GameTableDataManager{
		localizationDataMap: make(map[int]*localizationData),
	}
}

func (gtd *GameTableDataManager) changeReadTableType(_type gamedata.RTT_ReadTableType) gamedata.RTT_ReadTableType {
	switch _type {
	case gamedata.RTT_None:
		_type = gamedata.RTT_ReadStart
	case gamedata.RTT_ReadStart:
		_type = gamedata.RTT_ReadFinish
	}
	
	return _type
}

func (gtd *GameTableDataManager) allLoadData() error {
	if err := gtd.Load_EnglishWordData(); err != nil {
		return err
	}
	if err := gtd.Load_LocalizationData(); err != nil {
		return err
	}
	
	return nil
}

func (gtd *GameTableDataManager) RunGameTableDataServer(_callback func()) {
	if err := gtd.allLoadData(); err != nil {
		fmt.Printf("Error load table data : %v\n", err)
		return
	}
	
	if (_callback != nil) {
		_callback()
	}
}

func (gtd *GameTableDataManager) HttpHandle_load_englishworddata(res http.ResponseWriter, req *http.Request) {
	gtd.mtx.Lock()
	defer gtd.mtx.Unlock()
	
	data := make([]byte, gamedata.BufferSize)
	n, _ := req.Body.Read(data)
	var req_pack httpservermanager.Req_EnglishWordData
	json.Unmarshal([]byte(string(data[:n])), &req_pack)
	
	fmt.Printf("load_englishworddata req = %v\n", string(data[:n]))
	
	res_pack := &httpservermanager.Rsp_EnglishWordData{}
	
	bytess, err := gtd.GetJsonByEnlishWordSlice(req_pack.Idx)
	if err != nil {
		res_pack.Error = gamedata.EC_NotFoundTableInfo
		fmt.Println(err)
	} else {
		res_pack.Datas = string(bytess)
		fmt.Printf("load_englishworddata rsp : success.\n")
	}
	
	bytes, _ := json.Marshal(res_pack)
	
	renderer.Data(res, http.StatusOK, bytes)
}

func (gtd *GameTableDataManager) HttpHandle_load_localizationdata(res http.ResponseWriter, req *http.Request) {
	gtd.mtx.Lock()
	defer gtd.mtx.Unlock()
	
	data := make([]byte, gamedata.BufferSize)
	n, _ := req.Body.Read(data)
	var req_pack httpservermanager.Req_LocalizationData
	json.Unmarshal([]byte(string(data[:n])), &req_pack)
	
	fmt.Printf("load_localizationdata req = %v\n", string(data[:n]))
	
	res_pack := &httpservermanager.Rsp_LocalizationData{}
	
	bytess, err := gtd.GetJsonByLocalizationDataMap()
	if err != nil {
		res_pack.Error = gamedata.EC_NotFoundTableInfo
		fmt.Println(err)
	} else {
		res_pack.Datas = string(bytess)
		fmt.Printf("load_localizationdata rsp : success.\n")
	}
	
	bytes, _ := json.Marshal(res_pack)
	
	renderer.Data(res, http.StatusOK, bytes)
}
