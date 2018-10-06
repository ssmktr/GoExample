package gametabledata

import (
	"GoExample/gamedata"
	"fmt"
)

type GameTableDataManager struct {
	englishWordSlice []*englishWordData
}

func New() *GameTableDataManager {
	return &GameTableDataManager{
	}
}

func (gtd *GameTableDataManager) allLoadData() error {
	if err := gtd.Load_EnglishWordDataFile(); err != nil {
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

func (gtd *GameTableDataManager) changeReadTableType(_type gamedata.RTT_ReadTableType) gamedata.RTT_ReadTableType {
	switch _type {
	case gamedata.RTT_None:
		_type = gamedata.RTT_ReadStart
	case gamedata.RTT_ReadStart:
		_type = gamedata.RTT_ReadFinish
	}
	
	return _type
}

// func (gtd *GameTableDataManager) GetTypeByString(_type string) reflect.Type {
// 	switch _type {
// 	case "string":
// 		return reflect.Type(int)
// 	}
//
// 	return nil
// }
