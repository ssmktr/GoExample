package gametabledata

import (
	"testing"
)

func TestLoadLocalizationData(t *testing.T) {
	gtd := New()
	gtd.allLoadData()
	for key, data := range gtd.localizationDataMap {
		t.Logf("key : %v, data : %v\n", key, data)
	}
}
