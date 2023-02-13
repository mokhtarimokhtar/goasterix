package transform

import (
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"

	"github.com/mokhtarimokhtar/goasterix"
)

func TestCat021Model_ToJsonRecord(t *testing.T) {
	// Arrange
	input := "c51d3101432304 0001 0140 2bb73efa65ba 000001 384176 3adab9f5 00 02 00 08 cb 540d0d0d 0508f00162"
	output := []byte(`{"DataSourceIdentification":{"sac":0,"sic":1},"ServiceManagement":4,"EmitterCategory":"No ADS-B Emitter Category Information","TargetReportDescriptor":{"atp":"24-Bit ICAO address","arc":"25ft","rc":"Default","rab":"Report from target transponder","fx":{"dcr":"No differential correction","gbs":"Not set","sim":"Actual","tst":"Default","saa":"Capable","cl":"Report valid"}},"TimeOfMessageReceptionForPosition":28802.921875,"TimeOfMessageReceptionForPositionHighPrecision":{"FSI":"+0","TimeOfReception":0.9195999996736646},"TargetAddress":"000001","QualityIndicators":{},"PositionWGS84":{"latitude":61.47532332786,"longitude":-7.878698524580001},"MessageAmplitude":-53,"MOPSVersion":{"vns":"supported","vn":"ED102/DO-260","ltt":"1090 es"}}`)

	uap021 := uap.Cat021v10
	data, _ := util.HexStringToByte(input)
	rec := new(goasterix.Record)
	_, err := rec.Decode(data, uap021)

	cat021Model := new(Cat021Model)
	cat021Model.write(*rec)

	// Act
	recJson, _ := json.Marshal(cat021Model)

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

func Test_TargetReportDescriptors(t *testing.T) {
	// Arrange
	inputBytestream := []byte{0x2E}
	input := goasterix.Extended{
		Primary: inputBytestream,
	}
	output := TargetReportDescriptor{
		ATP: "Duplicate address",
		ARC: "100ft",
		RC:  "Range Check passed, CPR Validation pending",
		RAB: "Report from field monitor (fixed transponder)",
	}

	// Act
	res := targetReportDescriptor(input)

	// Assert
	if reflect.DeepEqual(output, res) {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_wgs84Coordinates_LowPrecision(t *testing.T) {
	// Arrange
	input := []byte{0x8a, 0x42, 0x02, 0x21, 0xff, 0x22}
	output := WGS84Coordinates{
		Latitude:  49.2139649391174,
		Longitude: 3.17801166325808,
	}
	epsilon := 0.1

	// Act
	res := wgs84Coordinates(input)

	// Assert
	if !checkEqualLatLong(res, output, epsilon) {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_wgs84Coordinates_HighPrecision(t *testing.T) {
	// Arrange
	input := []byte{0x11, 0x7f, 0x90, 0x06, 0x01, 0x21, 0x45, 0x0a}
	output := WGS84Coordinates{
		Latitude:  49.2139444872737,
		Longitude: 3.17801166325808,
	}
	epsilon := 0.1

	// Act
	res := wgs84Coordinates(input)

	// Assert
	if !checkEqualLatLong(res, output, epsilon) {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_BasicGeometricHeight(t *testing.T) {
	// Arrange
	input := [2]byte{0x09, 0x60}
	output := GeometricHeight{
		Height:      15000.0,
		GreaterThan: false,
	}

	// Act
	res := geometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_MaxGeometricHeight(t *testing.T) {
	// Arrange
	input := [2]byte{0x5D, 0xC0}
	output := GeometricHeight{
		Height:      150000.0,
		GreaterThan: false,
	}

	// Act
	res := geometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_MinGeometricHeight(t *testing.T) {
	// Arrange
	input := [2]byte{0xFF, 0x10}
	output := GeometricHeight{
		Height:      -1500.0,
		GreaterThan: false,
	}

	// Act
	res := geometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_GreaterThanGeometricHeight(t *testing.T) {
	// Arrange
	input := [2]byte{0x7F, 0xFF}
	output := GeometricHeight{
		Height:      204793.75,
		GreaterThan: true,
	}

	// Act
	res := geometricHeight(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}
