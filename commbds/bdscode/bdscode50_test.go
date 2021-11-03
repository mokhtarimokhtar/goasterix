package bdscode

import (
	"encoding/hex"
	"strings"
	"testing"
)

func Test_BDS_Code50_Decode_all_status_true(t *testing.T) {
	// Arrange
	input := "C0 FC 0F 8F 30 F6 0F"
	input = strings.ReplaceAll(input, " ", "")
	outputRollAngle := int8(-88)
	outputTrueTrackAngle := int8(-88)
	outputGroundSpeed := uint16(1144)
	outputTrackAngleRate := int8(-15)
	outputTrueAirSpeed := uint16(1054)

	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [7]byte
	copy(data[:], tmp) // convert slice to array for parameter function decode
	code50 := new(Code50)

	// Act
	err = code50.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if code50.RollAngleStatus != true {
		t.Errorf("FAIL: RollAngleStatus: %v; Expected: %v", code50.RollAngleStatus, true)
	} else {
		t.Logf("SUCCESS: RollAngleStatus: %v; Expected: %v", code50.RollAngleStatus, true)
	}
	if code50.RollAngle != outputRollAngle {
		t.Errorf("FAIL: RollAngle: %v; Expected: %v", code50.RollAngle, outputRollAngle)
	} else {
		t.Logf("SUCCESS: RollAngle: %v; Expected: %v", code50.RollAngle, outputRollAngle)
	}

	if code50.TrueTrackAngleStatus != true {
		t.Errorf("FAIL: TrueTrackAngleStatus: %v; Expected: %v", code50.TrueTrackAngleStatus, true)
	} else {
		t.Logf("SUCCESS: TrueTrackAngleStatus: %v; Expected: %v", code50.TrueTrackAngleStatus, true)
	}
	if code50.TrueTrackAngle != outputTrueTrackAngle {
		t.Errorf("FAIL: TrueTrackAngle: %v; Expected: %v", code50.TrueTrackAngle, outputTrueTrackAngle)
	} else {
		t.Logf("SUCCESS: TrueTrackAngle: %v; Expected: %v", code50.TrueTrackAngle, outputTrueTrackAngle)
	}

	if code50.GroundSpeedStatus != true {
		t.Errorf("FAIL: GroundSpeedStatus: %v; Expected: %v", code50.GroundSpeedStatus, true)
	} else {
		t.Logf("SUCCESS: GroundSpeedStatus: %v; Expected: %v", code50.GroundSpeedStatus, true)
	}
	if code50.GroundSpeed != outputGroundSpeed {
		t.Errorf("FAIL: GroundSpeed: %v; Expected: %v", code50.GroundSpeed, outputGroundSpeed)
	} else {
		t.Logf("SUCCESS: GroundSpeed: %v; Expected: %v", code50.GroundSpeed, outputGroundSpeed)
	}

	if code50.TrackAngleRateStatus != true {
		t.Errorf("FAIL: TrackAngleRateStatus: %v; Expected: %v", code50.TrackAngleRateStatus, true)
	} else {
		t.Logf("SUCCESS: TrackAngleRateStatus: %v; Expected: %v", code50.TrackAngleRateStatus, true)
	}
	if code50.TrackAngleRate != outputTrackAngleRate {
		t.Errorf("FAIL: TrackAngleRate: %v; Expected: %v", code50.TrackAngleRate, outputTrackAngleRate)
	} else {
		t.Logf("SUCCESS: TrackAngleRate: %v; Expected: %v", code50.TrackAngleRate, outputTrackAngleRate)
	}

	if code50.TrueAirSpeedStatus != true {
		t.Errorf("FAIL: TrueAirSpeedStatus: %v; Expected: %v", code50.TrueAirSpeedStatus, true)
	} else {
		t.Logf("SUCCESS: TrueAirSpeedStatus: %v; Expected: %v", code50.TrueAirSpeedStatus, true)
	}
	if code50.TrueAirSpeed != outputTrueAirSpeed {
		t.Errorf("FAIL: TrueAirSpeed: %v; Expected: %v", code50.TrueAirSpeed, outputTrueAirSpeed)
	} else {
		t.Logf("SUCCESS: TrueAirSpeed: %v; Expected: %v", code50.TrueAirSpeed, outputTrueAirSpeed)
	}
}

func Test_BDS_Code50_Decode_all_status_false(t *testing.T) {
	// Arrange
	input := "7F EF FE FF DF FB FF"
	input = strings.ReplaceAll(input, " ", "")
	outputRollAngle := int8(0)
	outputTrueTrackAngle := int8(0)
	outputGroundSpeed := uint16(0)
	outputTrackAngleRate := int8(0)
	outputTrueAirSpeed := uint16(0)

	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [7]byte
	copy(data[:], tmp) // convert slice to array for parameter function decode
	code50 := new(Code50)

	// Act
	err = code50.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if code50.RollAngleStatus != false {
		t.Errorf("FAIL: RollAngleStatus: %v; Expected: %v", code50.RollAngleStatus, false)
	} else {
		t.Logf("SUCCESS: RollAngleStatus: %v; Expected: %v", code50.RollAngleStatus, false)
	}
	if code50.RollAngle != outputRollAngle {
		t.Errorf("FAIL: RollAngle: %v; Expected: %v", code50.RollAngle, outputRollAngle)
	} else {
		t.Logf("SUCCESS: RollAngle: %v; Expected: %v", code50.RollAngle, outputRollAngle)
	}

	if code50.TrueTrackAngleStatus != false {
		t.Errorf("FAIL: TrueTrackAngleStatus: %v; Expected: %v", code50.TrueTrackAngleStatus, false)
	} else {
		t.Logf("SUCCESS: TrueTrackAngleStatus: %v; Expected: %v", code50.TrueTrackAngleStatus, false)
	}
	if code50.TrueTrackAngle != outputTrueTrackAngle {
		t.Errorf("FAIL: TrueTrackAngle: %v; Expected: %v", code50.TrueTrackAngle, outputTrueTrackAngle)
	} else {
		t.Logf("SUCCESS: TrueTrackAngle: %v; Expected: %v", code50.TrueTrackAngle, outputTrueTrackAngle)
	}

	if code50.GroundSpeedStatus != false {
		t.Errorf("FAIL: GroundSpeedStatus: %v; Expected: %v", code50.GroundSpeedStatus, false)
	} else {
		t.Logf("SUCCESS: GroundSpeedStatus: %v; Expected: %v", code50.GroundSpeedStatus, false)
	}
	if code50.GroundSpeed != outputGroundSpeed {
		t.Errorf("FAIL: GroundSpeed: %v; Expected: %v", code50.GroundSpeed, outputGroundSpeed)
	} else {
		t.Logf("SUCCESS: GroundSpeed: %v; Expected: %v", code50.GroundSpeed, outputGroundSpeed)
	}

	if code50.TrackAngleRateStatus != false {
		t.Errorf("FAIL: TrackAngleRateStatus: %v; Expected: %v", code50.TrackAngleRateStatus, false)
	} else {
		t.Logf("SUCCESS: TrackAngleRateStatus: %v; Expected: %v", code50.TrackAngleRateStatus, false)
	}
	if code50.TrackAngleRate != outputTrackAngleRate {
		t.Errorf("FAIL: TrackAngleRate: %v; Expected: %v", code50.TrackAngleRate, outputTrackAngleRate)
	} else {
		t.Logf("SUCCESS: TrackAngleRate: %v; Expected: %v", code50.TrackAngleRate, outputTrackAngleRate)
	}

	if code50.TrueAirSpeedStatus != false {
		t.Errorf("FAIL: TrueAirSpeedStatus: %v; Expected: %v", code50.TrueAirSpeedStatus, false)
	} else {
		t.Logf("SUCCESS: TrueAirSpeedStatus: %v; Expected: %v", code50.TrueAirSpeedStatus, false)
	}
	if code50.TrueAirSpeed != outputTrueAirSpeed {
		t.Errorf("FAIL: TrueAirSpeed: %v; Expected: %v", code50.TrueAirSpeed, outputTrueAirSpeed)
	} else {
		t.Logf("SUCCESS: TrueAirSpeed: %v; Expected: %v", code50.TrueAirSpeed, outputTrueAirSpeed)
	}
}
