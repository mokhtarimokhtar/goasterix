package transform

import (
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"reflect"
	"testing"
)

func TestCat004Model_ToJsonRecord(t *testing.T) {
	// Arrange
	input := "fdcb80 08a2 08 010882 6ae180 0000 08 0001 d1c0 41504d30303031 0001 0bc51ef7a55900f5 050370c30c40 00003039 ff50 ffd8a8 80 404cb3820820"

	output := []byte(`{"sourceIdentifier":{"sac":8,"sic":162},"messageType":{"code":"APM","desc":"approach_path_monitor"},"sdpsIdentifier":[{"sac":8,"sic":130}],"timeOfDay":54723}`)

	uap004 := uap.Cat004V112
	data, _ := goasterix.HexStringToByte(input)
	rec := goasterix.NewRecord()
	_, err := rec.Decode(data, uap004)

	model := new(Cat004Model)
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
