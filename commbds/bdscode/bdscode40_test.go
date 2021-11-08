package bdscode

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestBDSCode40Decode_AllStatusTrue(t *testing.T) {
	// Arrange
	input := "A3 2C 0F 30 A4 01 E7"
	input = strings.ReplaceAll(input, " ", "")
	outputMCPSelectAltitude := uint16(18000)
	outputFMSSelectAltitude := uint16(960)
	outputBarometricPressureSetting := uint16(213 + 800)
	outputVNAVMode := uint8(1)
	outputALTHOLDMode := uint8(1)
	outputAPPROACHMode := uint8(1)
	outputTargetAltSourceBits := uint8(3)
	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [7]byte
	copy(data[:], tmp) // convert slice to array for parameter function decode
	code40 := new(Code40)

	// Act
	err = code40.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if code40.MCPSelectAltitudeStatus != true {
		t.Errorf("FAIL: MCPSelectAltitudeStatus: %v; Expected: %v", code40.MCPSelectAltitudeStatus, true)
	} else {
		t.Logf("SUCCESS: MCPSelectAltitudeStatus: %v; Expected: %v", code40.MCPSelectAltitudeStatus, true)
	}
	if code40.MCPSelectAltitude != outputMCPSelectAltitude {
		t.Errorf("FAIL: MCPSelectAltitude: %v; Expected: %v", code40.MCPSelectAltitude, outputMCPSelectAltitude)
	} else {
		t.Logf("SUCCESS: MCPSelectAltitude: %v; Expected: %v", code40.MCPSelectAltitude, outputMCPSelectAltitude)
	}

	if code40.FMSSelectAltitudeStatus != true {
		t.Errorf("FAIL: FMSSelectAltitudeStatus: %v; Expected: %v", code40.FMSSelectAltitudeStatus, true)
	} else {
		t.Logf("SUCCESS: FMSSelectAltitudeStatus: %v; Expected: %v", code40.FMSSelectAltitudeStatus, true)
	}
	if code40.FMSSelectAltitude != outputFMSSelectAltitude {
		t.Errorf("FAIL: FMSSelectAltitude: %v; Expected: %v", code40.FMSSelectAltitude, outputFMSSelectAltitude)
	} else {
		t.Logf("SUCCESS: FMSSelectAltitude: %v; Expected: %v", code40.FMSSelectAltitude, outputFMSSelectAltitude)
	}

	if code40.BarometricPressureSettingStatus != true {
		t.Errorf("FAIL: BarometricPressureSettingStatus: %v; Expected: %v", code40.BarometricPressureSettingStatus, true)
	} else {
		t.Logf("SUCCESS: BarometricPressureSettingStatus: %v; Expected: %v", code40.BarometricPressureSettingStatus, true)
	}
	if code40.BarometricPressureSetting != outputBarometricPressureSetting {
		t.Errorf("FAIL: BarometricPressureSetting: %v; Expected: %v", code40.BarometricPressureSetting, outputBarometricPressureSetting)
	} else {
		t.Logf("SUCCESS: BarometricPressureSetting: %v; Expected: %v", code40.BarometricPressureSetting, outputBarometricPressureSetting)
	}

	if code40.MCPModeBitsStatus != true {
		t.Errorf("FAIL: MCPModeBitsStatus: %v; Expected: %v", code40.MCPModeBitsStatus, true)
	} else {
		t.Logf("SUCCESS: MCPModeBitsStatus: %v; Expected: %v", code40.MCPModeBitsStatus, true)
	}
	if code40.VNAVMode != outputVNAVMode {
		t.Errorf("FAIL: VNAVMode: %v; Expected: %v", code40.VNAVMode, outputVNAVMode)
	} else {
		t.Logf("SUCCESS: VNAVMode: %v; Expected: %v", code40.VNAVMode, outputVNAVMode)
	}
	if code40.ALTHOLDMode != outputALTHOLDMode {
		t.Errorf("FAIL: ALTHOLDMode: %v; Expected: %v", code40.ALTHOLDMode, outputALTHOLDMode)
	} else {
		t.Logf("SUCCESS: ALTHOLDMode: %v; Expected: %v", code40.ALTHOLDMode, outputALTHOLDMode)
	}
	if code40.APPROACHMode != outputAPPROACHMode {
		t.Errorf("FAIL: APPROACHMode: %v; Expected: %v", code40.APPROACHMode, outputAPPROACHMode)
	} else {
		t.Logf("SUCCESS: APPROACHMode: %v; Expected: %v", code40.APPROACHMode, outputAPPROACHMode)
	}

	if code40.TargetAltSourceBitsStatus != true {
		t.Errorf("FAIL: TargetAltSourceBitsStatus: %v; Expected: %v", code40.TargetAltSourceBitsStatus, true)
	} else {
		t.Logf("SUCCESS: TargetAltSourceBitsStatus: %v; Expected: %v", code40.TargetAltSourceBitsStatus, true)
	}
	if code40.TargetAltSourceBits != outputTargetAltSourceBits {
		t.Errorf("FAIL: TargetAltSourceBits: %v; Expected: %v", code40.TargetAltSourceBits, outputTargetAltSourceBits)
	} else {
		t.Logf("SUCCESS: TargetAltSourceBits: %v; Expected: %v", code40.TargetAltSourceBits, outputTargetAltSourceBits)
	}

}

func TestBDSCode40Decode_AllStatusFalse(t *testing.T) {
	// Arrange
	input := "7F FB FF DF FF FE FB"
	input = strings.ReplaceAll(input, " ", "")
	outputMCPSelectAltitude := uint16(0)
	outputFMSSelectAltitude := uint16(0)
	outputBarometricPressureSetting := uint16(0)
	outputVNAVMode := uint8(0)
	outputALTHOLDMode := uint8(0)
	outputAPPROACHMode := uint8(0)
	outputTargetAltSourceBits := uint8(0)
	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [7]byte
	copy(data[:], tmp) // convert slice to array for parameter function decode
	code40 := new(Code40)

	// Act
	err = code40.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if code40.MCPSelectAltitudeStatus != false {
		t.Errorf("FAIL: MCPSelectAltitudeStatus: %v; Expected: %v", code40.MCPSelectAltitudeStatus, false)
	} else {
		t.Logf("SUCCESS: MCPSelectAltitudeStatus: %v; Expected: %v", code40.MCPSelectAltitudeStatus, false)
	}
	if code40.MCPSelectAltitude != outputMCPSelectAltitude {
		t.Errorf("FAIL: MCPSelectAltitude: %v; Expected: %v", code40.MCPSelectAltitude, outputMCPSelectAltitude)
	} else {
		t.Logf("SUCCESS: MCPSelectAltitude: %v; Expected: %v", code40.MCPSelectAltitude, outputMCPSelectAltitude)
	}

	if code40.FMSSelectAltitudeStatus != false {
		t.Errorf("FAIL: FMSSelectAltitudeStatus: %v; Expected: %v", code40.FMSSelectAltitudeStatus, false)
	} else {
		t.Logf("SUCCESS: FMSSelectAltitudeStatus: %v; Expected: %v", code40.FMSSelectAltitudeStatus, false)
	}
	if code40.FMSSelectAltitude != outputFMSSelectAltitude {
		t.Errorf("FAIL: FMSSelectAltitude: %v; Expected: %v", code40.FMSSelectAltitude, outputFMSSelectAltitude)
	} else {
		t.Logf("SUCCESS: FMSSelectAltitude: %v; Expected: %v", code40.FMSSelectAltitude, outputFMSSelectAltitude)
	}

	if code40.BarometricPressureSettingStatus != false {
		t.Errorf("FAIL: BarometricPressureSettingStatus: %v; Expected: %v", code40.BarometricPressureSettingStatus, false)
	} else {
		t.Logf("SUCCESS: BarometricPressureSettingStatus: %v; Expected: %v", code40.BarometricPressureSettingStatus, false)
	}
	if code40.BarometricPressureSetting != outputBarometricPressureSetting {
		t.Errorf("FAIL: BarometricPressureSetting: %v; Expected: %v", code40.BarometricPressureSetting, outputBarometricPressureSetting)
	} else {
		t.Logf("SUCCESS: BarometricPressureSetting: %v; Expected: %v", code40.BarometricPressureSetting, outputBarometricPressureSetting)
	}

	if code40.MCPModeBitsStatus != false {
		t.Errorf("FAIL: MCPModeBitsStatus: %v; Expected: %v", code40.MCPModeBitsStatus, false)
	} else {
		t.Logf("SUCCESS: MCPModeBitsStatus: %v; Expected: %v", code40.MCPModeBitsStatus, false)
	}
	if code40.VNAVMode != outputVNAVMode {
		t.Errorf("FAIL: VNAVMode: %v; Expected: %v", code40.VNAVMode, outputVNAVMode)
	} else {
		t.Logf("SUCCESS: VNAVMode: %v; Expected: %v", code40.VNAVMode, outputVNAVMode)
	}
	if code40.ALTHOLDMode != outputALTHOLDMode {
		t.Errorf("FAIL: ALTHOLDMode: %v; Expected: %v", code40.ALTHOLDMode, outputALTHOLDMode)
	} else {
		t.Logf("SUCCESS: ALTHOLDMode: %v; Expected: %v", code40.ALTHOLDMode, outputALTHOLDMode)
	}
	if code40.APPROACHMode != outputAPPROACHMode {
		t.Errorf("FAIL: APPROACHMode: %v; Expected: %v", code40.APPROACHMode, outputAPPROACHMode)
	} else {
		t.Logf("SUCCESS: APPROACHMode: %v; Expected: %v", code40.APPROACHMode, outputAPPROACHMode)
	}

	if code40.TargetAltSourceBitsStatus != false {
		t.Errorf("FAIL: TargetAltSourceBitsStatus: %v; Expected: %v", code40.TargetAltSourceBitsStatus, false)
	} else {
		t.Logf("SUCCESS: TargetAltSourceBitsStatus: %v; Expected: %v", code40.TargetAltSourceBitsStatus, false)
	}
	if code40.TargetAltSourceBits != outputTargetAltSourceBits {
		t.Errorf("FAIL: TargetAltSourceBits: %v; Expected: %v", code40.TargetAltSourceBits, outputTargetAltSourceBits)
	} else {
		t.Logf("SUCCESS: TargetAltSourceBits: %v; Expected: %v", code40.TargetAltSourceBits, outputTargetAltSourceBits)
	}

}
