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

func (gtd *GameTableDataManager) Load_EnglishWordData() error {
	isReadData := gamedata.RTT_None
	fileName := gamedata.TableDataPath + "englishworddata.xlsx"
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
			
			checkRead := row.Cells[0].String()
			if checkRead == "@" {
				isReadData = gtd.changeReadTableType(isReadData)
				continue
			}
			if isReadData == gamedata.RTT_None {
				continue
			}
			if isReadData == gamedata.RTT_ReadFinish {
				break
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
			
			if englishData == "" {
				return fmt.Errorf("Error empty english data : %v", englishData)
			}
			
			if koreaData == "" {
				return fmt.Errorf("Error empty korea data : %v", koreaData)
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

func (gtd *GameTableDataManager) GetJsonByEnlishWordSlice(idx int) ([]byte, error) {
	englishWordSlice := []*englishWordData{}
	totalCount := 200
	startIdx := idx * totalCount
	endIdx := startIdx + totalCount
	
	for i := startIdx; i < endIdx; i++ {
		if len(gtd.englishWordSlice) <= i {
			return nil, fmt.Errorf("Error over idx GetJsonByEnlishWordSlice : %v", i)
		}
		englishWordSlice = append(englishWordSlice, gtd.englishWordSlice[i])
	}
	
	bytes, err := json.Marshal(englishWordSlice)
	if err != nil {
		return nil, fmt.Errorf("Error englishWordSlice : %v", err)
	}
	
	return bytes, err
}
