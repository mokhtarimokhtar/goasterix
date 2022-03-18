package transform

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

func TestCat063Model_ToJsonRecord(t *testing.T) {
	// Arrange
	//bff0090c79387308290000000012000000000000000000000000
	//input := "b0 0836 429b71 0801" //FSPEC : 10110000 => B0 f6 0836 429b71 0801 10940028200094008000

	input := "bff0 090c 79387308290000000012000000000000000000000000"

	output := []byte(`{"dataSourceIdentifier":{"sac":9,"sic":12},"timeOfMessage":62064.8984375,"sensorIdentifier":{"sac":8,"sic":41},"sensorConfigStatus":{"con":"operational","psr":"psr_go","ssr":"ssr_go","mds":"mode_s_go","ads":"ads_go","mlt":"mlt_go"},"timeStampingBias":0,"modeSRangeGainAndBias":{"srg":0.00018,"srb":0},"psrRangeGainAndBias":{"prg":0,"prb":0}}`)

	uap063 := uap.Cat063V16
	data, _ := goasterix.HexStringToByte(input)
	rec := goasterix.NewRecord()
	_, err := rec.Decode(data, uap063)

	model := new(Cat063Model)
	model.write(*rec)

	// Act
	recJson, _ := json.Marshal(model)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(recJson, output) == false {
		t.Errorf("FAIL: %s; \nExpected: %s", recJson, output)
	} else {
		t.Logf("SUCCESS: %s; Expected: %s", recJson, output)
	}
}
