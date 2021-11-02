package bdscode

import (
	"goasterix"
)

// Code50 Track and turn report
// RollAngle Range = [–90, + 90] degrees
// TrueTrackAngle Range = [–180, +180] degrees
// GroundSpeed Range = [0, 2 046] knots
// TrackAngleRate Range = [–16, +16] degrees/second
// TrueAirSpeed Range = [0, 2 046] knots
type Code50 struct {
	RollAngleStatus      bool   `json:"-"`
	RollAngle            int8   `json:"rollAngle,omitempty"`
	TrueTrackAngleStatus bool   `json:"-"`
	TrueTrackAngle       int8   `json:"trueTrackAngle,omitempty"`
	GroundSpeedStatus    bool   `json:"-"`
	GroundSpeed          uint16 `json:"groundSpeed,omitempty"`
	TrackAngleRateStatus bool   `json:"-"`
	TrackAngleRate       int8   `json:"trackAngleRate,omitempty"`
	TrueAirSpeedStatus   bool   `json:"-"`
	TrueAirSpeed         uint16 `json:"trueAirSpeed,omitempty"`
}

func (c *Code50) Decode(data [7]byte) (err error) {
	// Extract RollAngle
	raStatus := data[0] & 0x80 >> 7
	if raStatus == 1 {
		c.RollAngleStatus = true
		ra := uint16(data[0] & 0x7F) << 3 + uint16(data[1] & 0xE0) >> 5
		tmpRa :=  goasterix.TwoComplement16(10, ra)
		c.RollAngle = int8(float64(tmpRa) * 45/256)

	} else {
		c.RollAngleStatus = false
		c.RollAngle = 0
	}

	// TrueTrackAngle
	ttaStatus := data[1] & 0x10 >> 4
	if ttaStatus == 1 {
		c.TrueTrackAngleStatus = true
		tta := uint16(data[1] & 0x0F) << 7 + uint16(data[2] & 0xFE) >> 1
		tmpTta :=  goasterix.TwoComplement16(11, tta)
		c.TrueTrackAngle = int8(float64(tmpTta) * 90/512)

	} else {
		c.TrueTrackAngleStatus = false
		c.TrueTrackAngle = 0
	}

	// GroundSpeed
	gsStatus := data[2] & 0x01
	if gsStatus == 1 {
		c.GroundSpeedStatus = true
		gs := uint16(data[3] & 0xFF) << 2 + uint16(data[4] & 0xC0) >> 6
		c.GroundSpeed = gs * 2

	} else {
		c.GroundSpeedStatus = false
		c.GroundSpeed = 0
	}

	// TrackAngleRate
	tarStatus := data[4] & 0x20 >> 5
	if tarStatus == 1 {
		c.TrackAngleRateStatus = true
		tar := uint16(data[4] & 0x1F) << 5 + uint16(data[5] & 0xF8) >> 3
		tmpTar :=  goasterix.TwoComplement16(10, tar)
		c.TrackAngleRate = int8(float64(tmpTar) * 8/256)

	} else {
		c.TrackAngleRateStatus = false
		c.TrackAngleRate = 0
	}

	// TrueAirSpeed
	tasStatus := data[5] & 0x04 >> 2
	if tasStatus == 1 {
		c.TrueAirSpeedStatus = true
		tas := uint16(data[5] & 0x03) << 8 + uint16(data[6] & 0xFF)
		c.TrueAirSpeed = tas * 2

	} else {
		c.TrueAirSpeedStatus = false
		c.TrueAirSpeed = 0
	}

	return err
}