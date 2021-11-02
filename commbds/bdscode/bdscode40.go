package bdscode

// Code40 Selected vertical intention
// MCPSelectAltitude Range = [0, 65 520] feet
// FMSSelectAltitude Range = [0, 65 520] feet
// BarometricPressureSetting Range = [0, 410] mb
// TargetAltSourceBits
type Code40 struct {
	MCPSelectAltitudeStatus         bool   `json:"-"`
	MCPSelectAltitude               uint16 `json:"mcpSelectAltitude,omitempty"`
	FMSSelectAltitudeStatus         bool   `json:"-"`
	FMSSelectAltitude               uint16 `json:"fmsSelectAltitude,omitempty"`
	BarometricPressureSettingStatus bool   `json:"-"`
	BarometricPressureSetting       uint16 `json:"barometricPressureSetting,omitempty"`
	MCPModeBitsStatus               bool   `json:"-"`
	VNAVMode                        uint8  `json:"vnavMode,omitempty"`
	ALTHOLDMode                     uint8  `json:"altholdMode,omitempty"`
	APPROACHMode                    uint8  `json:"approachMode,omitempty"`
	TargetAltSourceBitsStatus       bool   `json:"-"`
	TargetAltSourceBits             uint8  `json:"targetAltSourceBits,omitempty"`
}

func (c *Code40) Decode(data [7]byte) (err error) {
	// Extract MCPSelectAltitude
	mcpStatus := data[0] & 0x80 >> 7
	if mcpStatus == 1 {
		c.MCPSelectAltitudeStatus = true
		mcp := uint16(data[0] & 0x7F) << 5 + uint16(data[1] & 0xF8) >> 3
		c.MCPSelectAltitude = mcp * 16
	} else {
		c.MCPSelectAltitudeStatus = false
		c.MCPSelectAltitude = 0
	}


	// Extract FMSSelectAltitude
	fmsStatus := data[1] & 0x04 >> 2
	if fmsStatus == 1 {
		c.FMSSelectAltitudeStatus = true
		fms := uint16(data[1] & 0x03) << 10 + uint16(data[2] & 0xFF) << 2 + uint16(data[3] & 0xC0) >> 6
		c.FMSSelectAltitude = fms * 16
	} else {
		c.FMSSelectAltitudeStatus = false
		c.FMSSelectAltitude = 0
	}


	// Extract BarometricPressureSetting
	bpsStatus := data[3] & 0x20 >> 5
	if bpsStatus == 1 {
		c.BarometricPressureSettingStatus = true
		bps := uint16(data[3] & 0x1F) << 7 + uint16(data[4] & 0xFE) >> 1
		c.BarometricPressureSetting = uint16(float64(bps) * 0.1) + 800
	} else {
		c.BarometricPressureSettingStatus = false
		c.BarometricPressureSetting = 0
	}


	// Extract ModeBits
	mbStatus := data[5] & 0x01
	if mbStatus == 1 {
		c.MCPModeBitsStatus = true
		c.VNAVMode = data[6] & 0x80 >> 7
		c.ALTHOLDMode = data[6] & 0x40 >> 6
		c.APPROACHMode = data[6] & 0x20 >> 5

	} else {
		c.MCPModeBitsStatus = false
		c.VNAVMode = 0
		c.ALTHOLDMode = 0
		c.APPROACHMode = 0
	}

	// Extract TargetAltSourceBits
	tasbStatus := data[6] & 0x04 >> 2
	if tasbStatus == 1 {
		c.TargetAltSourceBitsStatus = true
		c.TargetAltSourceBits = data[6] & 0x03

	} else {
		c.TargetAltSourceBitsStatus = false
		c.TargetAltSourceBits = 0
	}

	return nil
}
