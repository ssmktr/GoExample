package gametabledata

import (
	"GoExample/gamedata"
	"testing"
)

func TestLoadEnglishWordDataFile(t *testing.T) {
	gtd := New()
	gtd.allLoadData()
	for _, data := range gtd.englishWordSlice {
		t.Logf("english : %v, korea : %v\n", data.English, data.Korea)
	}
}

func TestIsInSliceData(t *testing.T) {
	gtd := New()
	isReadData := gamedata.RTT_None
	isReadData = gtd.changeReadTableType(isReadData)
	if isReadData == gamedata.RTT_ReadStart {
		t.Logf("right read type %v", isReadData)
	} else {
		t.Logf("wrong read type %v", isReadData)
	}
	
	isReadData = gtd.changeReadTableType(isReadData)
	if isReadData == gamedata.RTT_ReadFinish {
		t.Logf("right read type %v", isReadData)
	} else {
		t.Logf("wrong read type %v", isReadData)
	}
	
	isReadData = gtd.changeReadTableType(isReadData)
	if isReadData == gamedata.RTT_ReadFinish {
		t.Logf("right read type %v", isReadData)
	} else {
		t.Logf("wrong read type %v", isReadData)
	}
}