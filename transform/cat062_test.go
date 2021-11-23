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

func TestCat062Model_CalculatedAccelerationCartesian(t *testing.T) {
	// Arrange
	input := [2]byte{0x80, 0x10}
	output := Acceleration{
		Ax: -32,
		Ay: 4,
	}

	// Act
	res := calculatedAccelerationCartesian(input)

	// Assert
	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func TestCat062Model_TargetIdentification(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        [7]byte
		output       TargetIdent
	}
	dataset := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        [7]byte{0x00, 0x04, 0x64, 0xB1, 0xCB, 0x3D, 0x20},
			output: TargetIdent{
				Target: "AFR1234 ",
				STI:    "downlinked_target",
			},
		},
		{
			TestCaseName: "testcase 2",
			input:        [7]byte{0x40, 0x04, 0x64, 0xB1, 0xCB, 0x3D, 0x20},
			output: TargetIdent{
				Target: "AFR1234 ",
				STI:    "callsign_not_downlinked_target",
			},
		},
		{
			TestCaseName: "testcase 3",
			input:        [7]byte{0x80, 0x04, 0x64, 0xB1, 0xCB, 0x3D, 0x20},
			output: TargetIdent{
				Target: "AFR1234 ",
				STI:    "registration_not_downlinked_target",
			},
		},
		{
			TestCaseName: "testcase 4",
			input:        [7]byte{0xc0, 0x04, 0x64, 0xB1, 0xCB, 0x3D, 0x20},
			output: TargetIdent{
				Target: "AFR1234 ",
				STI:    "invalid",
			},
		},
	}

	for _, row := range dataset {
		// Arrange
		// Act
		s := targetIdentification(row.input)

		// Assert
		if s != row.output {
			t.Errorf("FAIL: %s - s = %s; Expected: %s", row.TestCaseName, s, row.output)
		} else {
			t.Logf("SUCCESS: s = %s; Expected: %s", s, row.output)
		}
	}
}

func TestCat062Model_TrackBarometricAltitude(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        [2]byte
		output       BarometricAltitude
	}
	dataset := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        [2]byte{0x00, 0x00},
			output: BarometricAltitude{
				QNH:      "no_qnh_correction_applied",
				Altitude: 0,
			},
		},
		{
			TestCaseName: "testcase 2",
			input:        [2]byte{0x80, 0xff},
			output: BarometricAltitude{
				QNH:      "qnh_correction_applied",
				Altitude: 63.75,
			},
		},
	}

	for _, row := range dataset {
		// Arrange
		// Act
		res := trackBarometricAltitude(row.input)

		// Assert
		if res != row.output {
			t.Errorf("FAIL: %s - res = %v; Expected: %v", row.TestCaseName, res, row.output)
		} else {
			t.Logf("SUCCESS: s = %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat062Model_Mode3ACode(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        [2]byte
		output       TrackMode3A
	}
	dataset := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        [2]byte{0x0F, 0xFF},
			output: TrackMode3A{
				V:      "code_validated",
				G:      "default",
				CH:     "no_change",
				Squawk: "7777",
			},
		},
		{
			TestCaseName: "testcase 2",
			input:        [2]byte{0xEF, 0xFF},
			output: TrackMode3A{
				V:      "code_not_validated",
				G:      "garbled_code",
				CH:     "changed",
				Squawk: "7777",
			},
		},
	}

	for _, row := range dataset {
		// Arrange
		// Act
		res := mode3ACode(row.input)

		// Assert
		if res != row.output {
			t.Errorf("FAIL: %s - res = %v; Expected: %v", row.TestCaseName, res, row.output)
		} else {
			t.Logf("SUCCESS: s = %v; Expected: %v", res, row.output)
		}
	}
}
