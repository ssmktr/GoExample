package gametabledata

import (
	"GoExample/gamedata"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
)

type localizationData struct {
	Index int
	Ko    string
	En    string
}

func (gtd *GameTableDataManager) Load_LocalizationData() error {
	isReadData := gamedata.RTT_None
	
	fileName := gamedata.TableDataPath + "localizationdata.xlsx"
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
			
			index := 0
			ko := ""
			en := ""
			for cellIdx, cell := range row.Cells {
				switch cellIdx {
				case 0:
					_index, err := cell.Int()
					if err != nil {
						return fmt.Errorf("Error localization data index type : %v", err)
					}
					fmt.Println(cell.String())
					index = _index
				case 1:
					ko = cell.String()
				case 2:
					en = cell.String()
				}
			}
			
			data := &localizationData{
				Index: index,
				Ko:    ko,
				En:    en,
			}
			
			if _, exist := gtd.localizationDataMap[index]; exist {
				return fmt.Errorf("Error is In localization data : %v", data)
			}
			
			gtd.localizationDataMap[index] = data
		}
	}
	
	return nil
}

func (gtd *GameTableDataManager) GetJsonByLocalizationDataMap() ([]byte, error) {
	bytes, err := json.Marshal(gtd.localizationDataMap)
	if err != nil {
		return nil, fmt.Errorf("Error localizationDataMap : %v", err)
	}
	
	return bytes, err
}
