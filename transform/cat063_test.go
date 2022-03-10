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
	input := "b0 0836 429b71 0801" //FSPEC : 10110000 => B0 f6 0836 429b71 0801 10940028200094008000
	output := []byte(`{"dataSourceIdentifier":{"sac":8,"sic":54},"timeOfMessage":34102.8828125,"sensorIdentifier":{"sac":8,"sic":1}}`)

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
