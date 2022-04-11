package transform

import (
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"
)

func TestCat004Model_MessageTypeCat004(t *testing.T) {
	// Arrange
	type dataTest struct {
		input  [1]byte
		output MsgType
	}
	dataset := []dataTest{
		{[1]byte{0x00}, MsgType{Code: msgTypeCode000, Desc: msgTypeDesc000}},
		{[1]byte{0x01}, MsgType{Code: msgTypeCode001, Desc: msgTypeDesc001}},
		{[1]byte{0x02}, MsgType{Code: msgTypeCode002, Desc: msgTypeDesc002}},
		{[1]byte{0x03}, MsgType{Code: msgTypeCode003, Desc: msgTypeDesc003}},
		{[1]byte{0x04}, MsgType{Code: msgTypeCode004, Desc: msgTypeDesc004}},
		{[1]byte{0x05}, MsgType{Code: msgTypeCode005, Desc: msgTypeDesc005}},
		{[1]byte{0x06}, MsgType{Code: msgTypeCode006, Desc: msgTypeDesc006}},
		{[1]byte{0x07}, MsgType{Code: msgTypeCode007, Desc: msgTypeDesc007}},
		{[1]byte{0x08}, MsgType{Code: msgTypeCode008, Desc: msgTypeDesc008}},
		{[1]byte{0x09}, MsgType{Code: msgTypeCode009, Desc: msgTypeDesc009}},
		{[1]byte{0x0a}, MsgType{Code: msgTypeCode010, Desc: msgTypeDesc010}},
		{[1]byte{0x0b}, MsgType{Code: msgTypeCode011, Desc: msgTypeDesc011}},
		{[1]byte{0x0c}, MsgType{Code: msgTypeCode012, Desc: msgTypeDesc012}},
		{[1]byte{0x0d}, MsgType{Code: msgTypeCode013, Desc: msgTypeDesc013}},
		{[1]byte{0x0e}, MsgType{Code: msgTypeCode014, Desc: msgTypeDesc014}},
		{[1]byte{0x0f}, MsgType{Code: msgTypeCode015, Desc: msgTypeDesc015}},
		{[1]byte{0x10}, MsgType{Code: msgTypeCode016, Desc: msgTypeDesc016}},
		{[1]byte{0x11}, MsgType{Code: msgTypeCode017, Desc: msgTypeDesc017}},
		{[1]byte{0x12}, MsgType{Code: msgTypeCode018, Desc: msgTypeDesc018}},
		{[1]byte{0x13}, MsgType{Code: msgTypeCode019, Desc: msgTypeDesc019}},
		{[1]byte{0x14}, MsgType{Code: msgTypeCode020, Desc: msgTypeDesc020}},
		{[1]byte{0x15}, MsgType{Code: msgTypeCode021, Desc: msgTypeDesc021}},
		{[1]byte{0x16}, MsgType{Code: msgTypeCode022, Desc: msgTypeDesc022}},
		{[1]byte{0x17}, MsgType{Code: msgTypeCode023, Desc: msgTypeDesc023}},
		{[1]byte{0x18}, MsgType{Code: msgTypeCode024, Desc: msgTypeDesc024}},
		{[1]byte{0x19}, MsgType{Code: msgTypeCode025, Desc: msgTypeDesc025}},
		{[1]byte{0x1a}, MsgType{Code: msgTypeCode026, Desc: msgTypeDesc026}},
		{[1]byte{0x1b}, MsgType{Code: msgTypeCode027, Desc: msgTypeDesc027}},
		{[1]byte{0x1c}, MsgType{Code: msgTypeCode028, Desc: msgTypeDesc028}},
		{[1]byte{0x1d}, MsgType{Code: msgTypeCode029, Desc: msgTypeDesc029}},
		{[1]byte{0x1e}, MsgType{Code: msgTypeCode030, Desc: msgTypeDesc030}},
		{[1]byte{0x1f}, MsgType{Code: msgTypeCode031, Desc: msgTypeDesc031}},
		{[1]byte{0x20}, MsgType{Code: msgTypeCode032, Desc: msgTypeDesc032}},
		{[1]byte{0x21}, MsgType{Code: msgTypeCode033, Desc: msgTypeDesc033}},
		{[1]byte{0x22}, MsgType{Code: msgTypeCode034, Desc: msgTypeDesc034}},
		{[1]byte{0x23}, MsgType{Code: msgTypeCode035, Desc: msgTypeDesc035}},
		{[1]byte{0x24}, MsgType{Code: msgTypeCode036, Desc: msgTypeDesc036}},
		{[1]byte{0x25}, MsgType{Code: msgTypeCode037, Desc: msgTypeDesc037}},
		{[1]byte{0x26}, MsgType{Code: msgTypeCode038, Desc: msgTypeDesc038}},
		{[1]byte{0x27}, MsgType{Code: msgTypeCode039, Desc: msgTypeDesc039}},
		{[1]byte{0x28}, MsgType{Code: msgTypeCode040, Desc: msgTypeDesc040}},
		{[1]byte{0x29}, MsgType{Code: msgTypeCode041, Desc: msgTypeDesc041}},
		{[1]byte{0x2a}, MsgType{Code: msgTypeCode042, Desc: msgTypeDesc042}},
		{[1]byte{0x2b}, MsgType{Code: msgTypeCode043, Desc: msgTypeDesc043}},
		{[1]byte{0x2c}, MsgType{Code: msgTypeCode044, Desc: msgTypeDesc044}},
		{[1]byte{0x61}, MsgType{Code: msgTypeCode097, Desc: msgTypeDesc097}},
		{[1]byte{0x62}, MsgType{Code: msgTypeCode098, Desc: msgTypeDesc098}},
		{[1]byte{0x63}, MsgType{Code: msgTypeCode099, Desc: msgTypeDesc099}},
	}
	for _, row := range dataset {
		// Act
		res := messageTypeCat004(row.input)

		// Assert
		if res != row.output {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat004Model_ToJsonRecord(t *testing.T) {
	// Arrange
	input := "fdcb80 08a2 08 010882 6ae180 0000 08 0001 d1c0 41504d30303031 0001 0bc51ef7a55900f5 050370c30c40 00003039 ff50 ffd8a8 80 404cb3820820"

	strOutput := `{
			"sourceIdentifier":{"sac":8,"sic":162},
			"messageType":{"code":"APM","desc":"approach_path_monitor"},
			"sdpsIdentifier":[{"sac":8,"sic":130}],
			"timeOfDay":54723,
			"alertIdentifier":0,
			"alertStatus":4,
			"trackNumber1":1,
			"verticalDeviation":-4400,
			"transversalDeviation":-5036
			}
			`
	output := []byte(util.CleanStringMultiline(strOutput))

	uap004 := uap.Cat004V112
	data, _ := util.HexStringToByte(input)
	rec := goasterix.NewRecord()
	_, err := rec.Decode(data, uap004)

	model := new(Cat004Model)
	model.write(*rec)

	// Act
	recJson, _ := json.Marshal(model)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v - Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v - Expected: %v", err, nil)
	}

	if reflect.DeepEqual(recJson, output) == false {
		t.Errorf("FAIL: %s - \nExpected: %s", recJson, output)
	} else {
		t.Logf("SUCCESS: %s - \nExpected: %s", recJson, output)
	}
}