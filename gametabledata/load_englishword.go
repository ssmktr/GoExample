package gametabledata

import (
	"GoExample/gamedata"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
)

type englishWordData struct {
	English string
	Korea   string
}

func (gtd *GameTableDataManager) Load_EnglishWordDataFile() error {
	isReadData := gamedata.RTT_None
	
	fileName := "../gametabledata/englishworddata.xlsx"
	xlsData, err := xlsx.OpenFile(fileName)
	if err != nil {
		return fmt.Errorf("Error ReadFile : %v\n", err)
	}
	
	for _, sheet := range xlsData.Sheets {
		if sheet.Name != "적용" {
			continue
		}
		
		for rowIdx, row := range sheet.Rows {
			if rowIdx == 0 {
				continue
			}
			
			englishData := ""
			koreaData := ""
			for cellIdx, cell := range row.Cells {
				switch cellIdx {
				case 0:
					englishData = cell.String()
				case 1:
					koreaData = cell.String()
				}
			}
			
			if englishData == "@" {
				isReadData = gtd.changeReadTableType(isReadData)
				continue
			}
			if isReadData == gamedata.RTT_None {
				continue
			}
			if isReadData == gamedata.RTT_ReadFinish {
				break
			}
			
			if englishData == "" {
				return fmt.Errorf(" empty english data : %v", englishData)
			}
			
			if koreaData == "" {
				return fmt.Errorf(" empty korea data : %v", koreaData)
			}
			
			data := &englishWordData{
				English: englishData,
				Korea:   koreaData,
			}
			
			if gtd.isInSliceData(data) {
				return fmt.Errorf("Error is In englishWordSlice : %v", data)
			}
			
			gtd.englishWordSlice = append(gtd.englishWordSlice, data)
		}
	}
	
	return nil
}

func (gtd *GameTableDataManager) isInSliceData(_data *englishWordData) bool {
	for _, data := range gtd.englishWordSlice {
		if _data.English == data.English {
			return true
		}
	}
	return false
}

func (gtd *GameTableDataManager) GetJsonByEnlishWordMap() ([]byte, error) {
	bytes, err := json.Marshal(gtd.englishWordSlice)
	if err != nil {
		return nil, fmt.Errorf("Error GetJsonByEnlishWordMap : %v", err)
	}
	
	return bytes, err
}
