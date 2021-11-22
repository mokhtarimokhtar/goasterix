package transform

import (
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"reflect"
	"testing"
)

func TestCat062Model_ToJsonRecord(t *testing.T) {
	// Arrange
	input := "bf5ffd0304 0900 01 532100 008e6f3e0017d096 1247f10b7086 fed3019a0fc8e301010c87304a04e072c34820e300820800eb003104b2190301487fa0ff0614ffffffffffff0493110101c006061414141400e0045b00e00182dc622931a410a800e00fc84010e001622b05010d01622902fea60177"
	output := []byte(`{"sourceIdentifier":{"sac":9,"sic":0},"serviceIdentification":1,"timeOfDay":42562,"trackPositionWGS84":{"latitude":50.07464289665222,"longitude":8.372386693954468},"cartesianXY":{"x":599032.5,"y":374851},"trackVelocity":{"vx":-75.25,"vy":102.5},"mode3ACode":{"v":"code_validated","g":"default","ch":"no_change","squawk":"7710"},"trackNumber":1202,"flightLevel":56,"geometricAltitude":6968.75,"barometricAltitude":{"qnh":"no_qnh_correction_applied","altitude":56},"rateOfClimbDescent":2412.5}`)

	uap062 := uap.Cat062V119
	data, _ := goasterix.HexStringToByte(input)
	rec := new(goasterix.Record)
	_, err := rec.Decode(data, uap062)

	cat062Model := new(Cat062Model)
	cat062Model.write(rec.Items)

	// Act
	recJson, _ := json.Marshal(cat062Model)

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
